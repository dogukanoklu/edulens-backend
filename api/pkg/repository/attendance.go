package repository

import (
	"api/pkg/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	AddAttendance(classID string, attendances []models.ReqAddAttendance) error
	GetAttendanceWithStudents(classID string, date int64) (*models.ResGetAttendance, error)
	UpdateAttendance(attendanceID string, updateAttendance []models.ReqUpdateAttendance) error
}

type GormAttendanceRepository struct {
	db *gorm.DB
}

func NewGormAttendanceRepository(db *gorm.DB) *GormAttendanceRepository {
	return &GormAttendanceRepository{db: db}
}

func (r *GormAttendanceRepository) AddAttendance(classID string, attendances []models.ReqAddAttendance) error {
	db := r.db

	attendance := models.Attendance{
		ClassID: classID,
	}

	result := db.Create(&attendance)

	if result.Error != nil {
		return result.Error
	}

	attendanceID := attendance.ID

	var valueStrings []string
	var valueArgs []interface{}

	for _, record := range attendances {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, attendanceID)
		valueArgs = append(valueArgs, record.UserID)
		valueArgs = append(valueArgs, record.IsPresent)
	}

	if len(valueStrings) > 0 {
		stmt := fmt.Sprintf("INSERT INTO attendance_details (attendance_id, user_id, is_present) VALUES %s", strings.Join(valueStrings, ","))

		result = db.Exec(stmt, valueArgs...)
		return result.Error
	}

	return nil
}

func (r *GormAttendanceRepository) GetAttendanceWithStudents(classID string, date int64) (*models.ResGetAttendance, error) {
	db := r.db

	var attendance models.ResGetAttendance

	// Query to get attendance data from the database using class_id
	err := db.Table("classes").
		Select("classes.id, classes.level, classes.branch, attendances.created_at").
		Joins("inner join students on classes.id = students.class_id").
		Joins("inner join attendances on classes.id = attendances.class_id").
		Joins("inner join attendance_details on attendances.id = attendance_details.attendance_id").
		Where("classes.id = ?", classID).
		Scan(&attendance).Error

	if err != nil {
		return nil, err
	}

	// Additional logic to fetch the students' details for the given date
	err = db.Table("students").
		Select("students.id, students.student_image, students.school_number, students.first_name, students.last_name, attendance_details.is_present").
		Joins("inner join attendance_details on students.id = attendance_details.user_id").
		Joins("inner join attendances on attendance_details.attendance_id = attendances.id").
		Where("attendances.class_id = ? AND DATE(FROM_UNIXTIME(attendances.created_at)) = DATE(FROM_UNIXTIME(?))", classID, date).
		Scan(&attendance.Students).Error

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

func (r *GormAttendanceRepository) UpdateAttendance(attendanceID string, updateAttendance []models.ReqUpdateAttendance) (error) {
	db := r.db

	for _, update := range updateAttendance {
		if err := db.Model(&models.AttendanceDetails{}).
		Where("attendance_id = ? AND user_id = ?", attendanceID, update.UserID).
		Update("is_present", update.IsPresent).Error; err != nil {
			return err
		}
	}

	return nil
}
