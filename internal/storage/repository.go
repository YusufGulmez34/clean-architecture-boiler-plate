package storage

import "boiler-plate/models"

type IBaseRepository[TModel any] interface {
	Create(model *TModel) error
	FindAll() ([]*TModel, error)
	FindByID(id uint) (*TModel, error)
	Update(model *TModel) error
	Delete(id uint) error
}

type IProductRepository interface {
	IBaseRepository[models.Product]
}

type ITokenRepository interface {
	IBaseRepository[models.TokenDetail]
	GetByUserID(userID uint) (models.TokenDetail, error)
	GetByUUID(uuid string) (models.TokenDetail, error)
}

type IUserRepository interface {
	IBaseRepository[models.User]
	FindByUsername(username string) (*models.User, error)
}
