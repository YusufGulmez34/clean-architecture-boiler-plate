package gormproductrepo

import (
	"boiler-plate/internal/storage"
	"boiler-plate/internal/storage/gormrepository"
	"boiler-plate/models"

	"gorm.io/gorm"
)

type GormProductRepository struct {
	gormrepository.GormBaseRepository[models.Product]
}

func NewGormProductRepository(db *gorm.DB) storage.IProductRepository {
	db.AutoMigrate(&models.Product{})
	productRepo := &GormProductRepository{}
	productRepo.NewGormBaseRepository(db)
	return productRepo
}
