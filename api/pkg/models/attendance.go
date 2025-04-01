package models

import "gorm.io/gorm"

// Database Models

type Student struct {
	ID             string `json:"id"`
	ClassID        int64  `json:"classID"`
	ProfilePicture string `json:"profilePicture"`
	SchoolNumber   int64  `json:"schoolNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	CreatedAt      int    `json:"createdAt"`
}

type Attendance struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"` // Eğer ID auto-increment integer ise
	ClassID   string `json:"classID"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdAt"` // GORM otomatik oluşturur
}

type AttendanceDetails struct {
	ID           string `json:"id"`
	AttendanceID string `json:"attendanceID"`
	UserID       string `json:"userID"`
	IsPresent    bool   `json:"isPresent"`
}

// Request Models

type ReqAddAttendance struct {
	UserID    string `json:"userID"`
	IsPresent bool   `json:"isPresent"`
}

type ReqUpdateAttendance struct {
	UserID    string `json:"userID"`
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
	ID             string `json:"id"`
	ProfilePicture string `json:"profilePicture"`
	SchoolNumber   int64  `json:"schoolNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
}

func (student *ResStudent) AfterFind(tx *gorm.DB) (err error) {
	student.ProfilePicture = "http://localhost:8000/images/" + student.ProfilePicture
	return nil
}

// Database Add

type AddAttendances struct {
	ID          string `gorm:"primaryKey;autoIncrement"`
	Attendances []ReqAddAttendance
}
