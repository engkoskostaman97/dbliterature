package repositories

import (
	"literatur/models"

	"gorm.io/gorm"
)

type CollectionRepository interface {
	FindCollections() ([]models.Collection, error)
	GetCollection(ID int) (models.Collection, error)
	CreateCollection(Collection models.Collection) (models.Collection, error)
	UpdateCollection(Collection []models.Collection) ([]models.Collection, error)
	UpdateeCollection(Collection models.Collection) (models.Collection, error)
	DeleteCollection(Collection models.Collection) (models.Collection, error)

}

func RepositoryCollection(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCollections() ([]models.Collection, error) {
	var collections []models.Collection
	err := r.db.Preload("User").Preload("Literatur").Preload("Literatur.User").Find(&collections).Error

	return collections, err
}

func (r *repository) GetCollection(ID int) (models.Collection, error) {
	var collection models.Collection
	err := r.db.Preload("Literatur").Preload("User").Preload("Literatur.User").First(&collection, ID).Error

	return collection, err
}

func (r *repository) CreateCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Preload("Literatur").Create(&collection).Error

	return collection, err
}

func (r *repository) UpdateCollection(collection []models.Collection) ([]models.Collection, error) {
	err := r.db.Save(&collection).Error

	return collection, err
}

func (r *repository) UpdateeCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Save(&collection).Error

	return collection, err
}

func (r *repository) DeleteCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Delete(&collection).Error

	return collection, err
}


func (r *repository) FindLiteraturID([]int) ([]models.Literatur, error) {
	var literatur []models.Literatur
	err := r.db.Find(&literatur).Error

	return literatur, err
}

