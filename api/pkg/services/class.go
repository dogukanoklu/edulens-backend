package services

import (
	"api/pkg/models"
	"api/pkg/repository"
)

type ClassService struct {
	repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) GetClasses() (*[]models.ResClasses, error) {
	return s.repo.GetClasses()
}
