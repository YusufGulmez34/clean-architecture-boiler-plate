package gormuserrepo

import (
	"boiler-plate/internal/storage"
	"boiler-plate/internal/storage/gormrepository"
	"boiler-plate/models"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	gormrepository.GormBaseRepository[models.User]
}

func NewGormUserRepository(db *gorm.DB) storage.IUserRepository {
	db.AutoMigrate(&models.User{})
	userRepository := &GormUserRepository{}
	userRepository.NewGormBaseRepository(db)
	return userRepository
}

func (u *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.DB().Where("username=?", username).First(&user).Error
	return &user, err
}
