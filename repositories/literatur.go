package repositories

import (
	"literatur/models"

	"gorm.io/gorm"
)

type LiteraturRepository interface {
	FindLiteraturs() ([]models.Literatur, error)
	GetLiteratur(ID int) (models.Literatur, error)
	CreateLiteratur(literatur models.Literatur) (models.Literatur, error)
	
}

func RepositoryLiteratur(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLiteraturs() ([]models.Literatur, error) {
	var literaturs []models.Literatur
	err := r.db.Preload("User").Find(&literaturs).Error

	return literaturs, err
}

func (r *repository) GetLiteratur(ID int) (models.Literatur, error) {
	var literatur models.Literatur
	err := r.db.Preload("User").First(&literatur, ID).Error

	return literatur, err
}

func (r *repository) CreateLiteratur(literatur models.Literatur) (models.Literatur, error) {
	err := r.db.Create(&literatur).Error

	return literatur, err
}

