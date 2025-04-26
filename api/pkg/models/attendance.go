package models

import "gorm.io/gorm"

// Database Models
type Student struct {
	ID           string `json:"id"`
	ClassID      int64  `json:"classID"`
	StudentImage string `json:"studentImage"`
	SchoolNumber int64  `json:"schoolNumber"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CreatedAt    int    `json:"createdAt"`
}

type Attendance struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ClassID   string `json:"classID"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdAt"`
}

type AttendanceDetails struct {
	ID           string `json:"id"`
	AttendanceID string `json:"attendanceID"`
	StudentID    string `json:"studentID"`
	IsPresent    bool   `json:"isPresent"`
}

// Request Models
type ReqAddAttendance struct {
	StudentID string `json:"studentID"`
	IsPresent bool   `json:"isPresent"`
}

type ReqUpdateAttendance struct {
	StudentID string `json:"studentID"`
	IsPresent bool   `json:"isPresent"`
}

// Response Models
type ResGetAttendance struct {
	ID        string       `json:"id"`
	Level     int          `json:"level"`
	Branch    string       `json:"branch"`
	Students  []ResStudent `json:"students"`
	CreatedAt int64        `json:"createdAt"`
}

type ResStudent struct {
<<<<<<< HEAD
	ID           string `json:"id"`
	StudentImage string `json:"studentImage"`
	SchoolNumber int64  `json:"schoolNumber"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
=======
	ID             string `json:"id"`
	StudentImage string `json:"studentImage"`
	SchoolNumber   int64  `json:"schoolNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
>>>>>>> 0ad2f7cd5d8e00c3deb2064cc3a8e31d939b9f66
}

func (student *ResStudent) AfterFind(tx *gorm.DB) (err error) {
	student.StudentImage = "http://localhost:8000/images/" + student.StudentImage
	return nil
}

// Database Add
type AddAttendances struct {
	ID          string `gorm:"primaryKey;autoIncrement"`
	Attendances []ReqAddAttendance
}
