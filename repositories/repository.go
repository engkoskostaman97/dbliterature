package repositories

import (
	"gorm.io/gorm"
	"literatur/models"
)

type repository struct {
	db *gorm.DB
}

// FindCollectionsTransaction implements CollectionRepository
func (*repository) FindCollectionsTransaction(TrxID int) ([]models.Collection, error) {
	panic("unimplemented")
}
