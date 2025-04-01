package services

import (
	"api/pkg/models"
	"api/pkg/repository"
)

type AttendanceService struct {
	repo repository.AttendanceRepository
}

func NewAttendanceService(repo repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{repo: repo}
}

func (s *AttendanceService) AddAttendance(classID string, attendances []models.ReqAddAttendance) error {
	return s.repo.AddAttendance(classID, attendances)
}

func (s *AttendanceService) GetAttendanceWithStudents(classID string, date int64) (*models.ResGetAttendance, error) {
	return s.repo.GetAttendanceWithStudents(classID, date)
}

func (s *AttendanceService) UpdateAttendance(attendanceID string, updateAttendance []models.ReqUpdateAttendance) error {
	return s.repo.UpdateAttendance(attendanceID, updateAttendance)
}