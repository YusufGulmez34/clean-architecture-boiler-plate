package services

import (
	"boiler-plate/internal/storage"
	"boiler-plate/models"
)

type IBaseService[TModel any, TRequest any, TResponse any] interface {
	Storage() storage.IStorage
	Add(modelDTO TRequest) error
	GetAll() ([]*TResponse, error)
	GetByID(id uint) (*TResponse, error)
	Delete(id uint) error
	Update(modelDTO TRequest, id uint) error
}

type IProductService interface {
	IBaseService[models.Product, models.ProductRequestDTO, models.ProductResponsDTO]
}

type IAuthService interface {
	Login(userLoginDTO models.UserLoginDTO) (tokenString string, err error)
	Register(userRegisterDTO models.UserRegisterDTO) error
}
