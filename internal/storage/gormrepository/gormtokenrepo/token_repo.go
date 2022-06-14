package gormtokenrepo

import (
	"boiler-plate/internal/storage"
	"boiler-plate/internal/storage/gormrepository"
	"boiler-plate/models"

	"gorm.io/gorm"
)

type GormTokenRepository struct {
	gormrepository.GormBaseRepository[models.TokenDetail]
}

func NewGormTokenRepository(db *gorm.DB) storage.ITokenRepository {
	db.AutoMigrate(&models.TokenDetail{})
	tokenRepository := &GormTokenRepository{}
	tokenRepository.NewGormBaseRepository(db)
	return tokenRepository
}

func (t *GormTokenRepository) GetByUserID(userID uint) (models.TokenDetail, error) {
	var tokenDetail models.TokenDetail
	err := t.DB().Where("user_id = ?", userID).First(&tokenDetail).Error
	return tokenDetail, err
}

func (t *GormTokenRepository) GetByUUID(uuid string) (models.TokenDetail, error) {
	var tokenDetail models.TokenDetail
	err := t.DB().Where("uuid = ?", uuid).First(&tokenDetail).Error
	return tokenDetail, err
}
