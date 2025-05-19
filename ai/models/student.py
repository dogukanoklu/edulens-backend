from dataclasses import dataclass

@dataclass
class Student:
    id: str
    school_number: int
    first_name: str
    last_name: str
    level: int
    branch: str
    student_image: int = 1

    def to_dict(self):
        return {
            "id": self.id,
            "schoolNumber": self.school_number,
            "studentImage": f"http://localhost:8000/image/{self.student_image}",
            "firstName": self.first_name,
            "lastName": self.last_name,
            "level": self.level,
            "branch": self.branch
        }
