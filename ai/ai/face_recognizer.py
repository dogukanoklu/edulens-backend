import cv2
import face_recognition
import numpy as np
from utils.database import Database
from models.student import Student

class FaceRecognizer: 
    def __init__(self, db_config):
        self.db = Database(**db_config)
        self.known_encodings, self.student_infos = self.db.load_students()
        print(f"Loaded Face Codes: {len(self.known_encodings)} piece")
        print(f"Loaded Student Infos: {len(self.student_infos)} piece")
        
    def recognize_faces(self, image_bgr) -> Student | None:
        print("Scanning for faces...")
        small_frame = cv2.resize(image_bgr, (0, 0), fx=0.5, fy=0.5)
        rgb_small_frame = small_frame[:, :, ::-1]

        face_locations = face_recognition.face_locations(rgb_small_frame)
        print(f"Face locations: {face_locations}")

        face_encodings = face_recognition.face_encodings(rgb_small_frame, face_locations)
        print(f"Face encodings found: {len(face_encodings)}")

        if face_encodings:
            face_encoding = face_encodings[0]

            matches = face_recognition.compare_faces(self.known_encodings, face_encoding)
            face_distances = face_recognition.face_distance(self.known_encodings, face_encoding)
            print(f"Matches: {matches}")
            print(f"Face distances: {face_distances}")

            best_match_index = np.argmin(face_distances) if face_distances.size > 0 else None

            if best_match_index is not None and matches[best_match_index]:
                student = self.student_infos[best_match_index]
                print(f"Recognized student: {student['first_name']} {student['last_name']} ({student['school_number']})")

                return Student(
                    id=student.get("id"),
                    school_number=student.get("school_number"),
                    first_name=student.get("first_name"),
                    last_name=student.get("last_name"),
                    level=student.get("level"),
                    branch=student.get("branch"),
                    student_image=student.get("student_image", 1)
                )
            else:
                print("Face match not found.")
                return None
        else:
            print("No faces found.")
            return None
