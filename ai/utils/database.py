import mysql.connector
import face_recognition
import pickle
from pathlib import Path
import uuid
import time

class Database:
    def __init__(self, host, user, password, database):
        self.config = {
            'host': host,
            'user': user,
            'password': password,
            'database': database
        }

        # Path to the cache file
        self.cache_file = Path(__file__).resolve().parent.parent / "data" / "student_cache.pkl"

    def connect(self):
        """Establishes a connection to the MySQL database."""
        return mysql.connector.connect(**self.config)

    def clear_cache(self):
        """Deletes the cached face encodings and student info."""
        if self.cache_file.exists():
            self.cache_file.unlink()
            print("[INFO] Cache cleared.")

    def load_students(self, force_reload=False):
        """
        Loads student data and face encodings.

        If a cache exists and force_reload is False, data will be loaded from cache.
        Otherwise, face encodings are re-computed from student images.

        :param force_reload: Set to True to ignore cache and reload everything.
        :return: Tuple of (known_encodings, student_infos)
        """
        if self.cache_file.exists() and not force_reload:
            print("[INFO] Loading face data from cache...")
            try:
                with open(self.cache_file, "rb") as f:
                    known_encodings, student_infos = pickle.load(f)
                return known_encodings, student_infos
            except Exception as e:
                print(f"[WARNING] Failed to load cache, reloading data... ({e})")

        # Fetch student info from the database
        conn = self.connect()
        cursor = conn.cursor()
        cursor.execute("""
            SELECT 
                students.id,
                students.first_name,
                students.last_name,
                students.school_number,
                classes.level,
                classes.branch
            FROM students
            JOIN classes ON students.class_id = classes.id;
        """)
        students = cursor.fetchall()
        cursor.close()
        conn.close()

        known_encodings = []
        student_infos = []

        # Path to the student image folder (../data/student_images)
        image_folder = Path(__file__).resolve().parent.parent / "data" / "student_images"

        for student in students:
            student_id, first_name, last_name, school_number, level, branch = student
            image_path = image_folder / f"{school_number}.jpg"

            if not image_path.exists():
                print(f"[WARNING] Photo not found: {image_path}")
                continue

            try:
                image = face_recognition.load_image_file(image_path)
                encodings = face_recognition.face_encodings(image)

                if encodings:
                    known_encodings.append(encodings[0])
                    student_infos.append({
                        "id": student_id,
                        "first_name": first_name,
                        "last_name": last_name,
                        "school_number": school_number,
                        "level": level,
                        "branch": branch,
                        "student_image": school_number  # Image identifier
                    })
                else:
                    print(f"[WARNING] No face found in image: {image_path}")
            except Exception as e:
                print(f"[ERROR] Could not process image {image_path} -> {e}")

        # Save to cache
        try:
            with open(self.cache_file, "wb") as f:
                pickle.dump((known_encodings, student_infos), f)
            print("[INFO] Face data cached successfully.")
        except Exception as e:
            print(f"[ERROR] Failed to write cache file -> {e}")

        return known_encodings, student_infos
    
    def create_attendance_if_needed_for_student(self, student_id):
        conn = self.connect()
        cursor = conn.cursor()

        # 1. Get the student's class_id
        cursor.execute("SELECT class_id FROM students WHERE id = %s", (student_id,))
        result = cursor.fetchone()
        if not result:
            print("[ERROR] Student not found.")
            cursor.close()
            conn.close()
            return False

        class_id = result[0]

        # 2. Check if attendance already exists for the class (without date check)
        cursor.execute("""
            SELECT id FROM attendances
            WHERE class_id = %s
            LIMIT 1
        """, (class_id,))
        result = cursor.fetchone()

        if result:
            print("[INFO] Attendance already exists. No action taken.")
            cursor.close()
            conn.close()
            return False

        # 3. Create a new attendance session (without created_at)
        attendance_id = str(uuid.uuid4())
        cursor.execute("""
            INSERT INTO attendances (id, class_id)
            VALUES (%s, %s)
        """, (attendance_id, class_id))
        print(f"[INFO] Created new attendance session: {attendance_id}")

        # 4. Get all students in the class
        cursor.execute("SELECT id FROM students WHERE class_id = %s", (class_id,))
        student_ids = cursor.fetchall()

        # 5. Insert attendance details for each student
        for (s_id,) in student_ids:
            detail_id = str(uuid.uuid4())
            is_present = (s_id == student_id)

            cursor.execute("""
                INSERT INTO attendance_details (id, attendance_id, student_id, is_present)
                VALUES (%s, %s, %s, %s)
            """, (detail_id, attendance_id, s_id, is_present))

        conn.commit()
        cursor.close()
        conn.close()

        print("[SUCCESS] Attendance and details successfully recorded.")
        return True
