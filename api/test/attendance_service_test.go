package test

import (
	"api/pkg/models"
	"api/pkg/services"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAttendanceRepository struct {
	mock.Mock
}

func (m *MockAttendanceRepository) AddAttendance(classID string, attendances []models.ReqAddAttendance) error {
	args := m.Called(classID, attendances)
	return args.Error(0)
}

func (m *MockAttendanceRepository) GetAttendanceWithStudents(classID string, date int64) (*models.ResGetAttendance, error) {
	args := m.Called(classID, date)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ResGetAttendance), args.Error(1)
}

func setupAddAttendanceTest(tb testing.TB) (*MockAttendanceRepository, *services.AttendanceService, string, []models.ReqAddAttendance) {
	mockRepo := new(MockAttendanceRepository)
	classID := "196d3e7e-0e59-11f0-9064-c85b764bb385"
	attendances := []models.ReqAddAttendance{
		{UserID: "f7fe49f3-0e59-11f0-9064-c85b764bb385", IsPresent: true},
		{UserID: "f7fb5ea7-0e59-11f0-9064-c85b764bb385", IsPresent: false},
	}
	attendanceService := services.NewAttendanceService(mockRepo)
	return mockRepo, attendanceService, classID, attendances
}

func setupAttendanceTest(t *testing.T) (*MockAttendanceRepository, *services.AttendanceService, string, int64, *models.ResGetAttendance) {
	mockRepo := new(MockAttendanceRepository)
	classID := "196d3e7e-0e59-11f0-9064-c85b764bb385"
	date := time.Now().Unix()
	expectedAttendance := &models.ResGetAttendance{}
	mockRepo.On("GetAttendanceWithStudents", classID, date).Return(expectedAttendance, nil)
	attendanceService := services.NewAttendanceService(mockRepo)
	return mockRepo, attendanceService, classID, date, expectedAttendance
}


func TestAddAttendance_Success(t *testing.T) {
	mockRepo, attendanceService, classID, attendances := setupAddAttendanceTest(t)

	mockRepo.On("AddAttendance", classID, attendances).Return(nil)

	err := attendanceService.AddAttendance(classID, attendances)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddAttendance_Failure(t *testing.T) {
	mockRepo, attendanceService, classID, attendances := setupAddAttendanceTest(t)
	expectedError := errors.New("database error")

	mockRepo.On("AddAttendance", classID, attendances).Return(expectedError)

	err := attendanceService.AddAttendance(classID, attendances)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func BenchmarkAddAttendance(b *testing.B) {
	mockRepo, attendanceService, classID, attendances := setupAddAttendanceTest(b)

	mockRepo.On("AddAttendance", classID, attendances).Return(nil)

	for i := 0; i < b.N; i++ {
		err := attendanceService.AddAttendance(classID, attendances)
		if err != nil {
			b.Fatalf("Test sırasında hata oluştu: %v", err)
		}
	}
	mockRepo.AssertExpectations(b)
}

func TestGetAttendanceWithStudents_Success(t *testing.T) {
	mockRepo, attendanceService, classID, date, expectedAttendance := setupAttendanceTest(t)

	result, err := attendanceService.GetAttendanceWithStudents(classID, date)

	assert.NoError(t, err)
	assert.Equal(t, expectedAttendance, result)
	mockRepo.AssertExpectations(t)
}

func BenchmarkGetAttendanceWithStudents(b *testing.B) {
	mockRepo := new(MockAttendanceRepository)
	classID := "196d3e7e-0e59-11f0-9064-c85b764bb385"
	date := time.Now().Unix()
	expectedAttendance := &models.ResGetAttendance{}
	mockRepo.On("GetAttendanceWithStudents", classID, date).Return(expectedAttendance, nil)
	attendanceService := services.NewAttendanceService(mockRepo)

	for i := 0; i < b.N; i++ {
		_, err := attendanceService.GetAttendanceWithStudents(classID, date)
		if err != nil {
			b.Fatalf("Test sırasında hata oluştu: %v", err)
		}
	}
	mockRepo.AssertExpectations(b)
}