package repository

import (
	"api/pkg/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	GetClasses() (*[]models.ResClasses, error)
}

type GormClassRepository struct {
	db *gorm.DB
}

func NewGormClassRepository(db *gorm.DB) *GormClassRepository {
	return &GormClassRepository{db: db}
}

func (r *GormClassRepository) GetClasses() (*[]models.ResClasses, error) {
	db := r.db

	var classes []models.ResClasses

	// Query to get attendance data from the database using class_id
	err := db.Table("classes").
		Select("id, level, branch").
		Scan(&classes).Error

	if err != nil {
		return nil, err
	}

	return &classes, nil
}
