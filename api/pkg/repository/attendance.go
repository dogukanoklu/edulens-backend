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
		valueArgs = append(valueArgs, record.StudentID)
		valueArgs = append(valueArgs, record.IsPresent)
	}

	if len(valueStrings) > 0 {
		stmt := fmt.Sprintf("INSERT INTO attendance_details (attendance_id, student_id, is_present) VALUES %s", strings.Join(valueStrings, ","))

		result = db.Exec(stmt, valueArgs...)
		return result.Error
	}

	return nil
}

func (r *GormAttendanceRepository) GetAttendanceWithStudents(classID string, date int64) (*models.ResGetAttendance, error) {
	db := r.db

	// Start of the day and end of the day in seconds
	startOfDay := date / 1000 * 1000
	endOfDay := startOfDay + 86_399_999

	type ClassAttendanceData struct {
		ID        string `gorm:"column:id"`
		Level     int    `gorm:"column:level"`
		Branch    string `gorm:"column:branch"`
		CreatedAt int64  `gorm:"column:created_at"`
	}

	var tempData ClassAttendanceData

	// Fetch the latest attendance data for the class
	err := db.Table("classes").
		Select("attendances.id, classes.level, classes.branch, attendances.created_at").
		Joins("inner join attendances on classes.id = attendances.class_id").
		Where("classes.id = ?", classID).
		Order("attendances.created_at DESC").
		Limit(1).
		Scan(&tempData).Error

	if err != nil {
		return nil, err
	}

	attendance := &models.ResGetAttendance{
		ID:        tempData.ID,
		Level:     tempData.Level,
		Branch:    tempData.Branch,
		CreatedAt: tempData.CreatedAt,
	}

	// Fetch students based on the attendance created_at range for that day
	err = db.Table("students").
		Select("students.id, students.student_image, students.school_number, students.first_name, students.last_name, attendance_details.is_present").
		Joins("inner join attendance_details on students.id = attendance_details.student_id").
		Joins("inner join attendances on attendance_details.attendance_id = attendances.id").
		Where("attendances.class_id = ? AND attendances.created_at BETWEEN ? AND ?", classID, startOfDay, endOfDay).
		Find(&attendance.Students).Error

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (r *GormAttendanceRepository) UpdateAttendance(attendanceID string, updateAttendance []models.ReqUpdateAttendance) error {
	db := r.db

	for _, update := range updateAttendance {
		if err := db.Model(&models.AttendanceDetails{}).
			Where("attendance_id = ? AND student_id = ?", attendanceID, update.StudentID).
			Update("is_present", update.IsPresent).Error; err != nil {
			return err
		}
	}

	return nil
}
