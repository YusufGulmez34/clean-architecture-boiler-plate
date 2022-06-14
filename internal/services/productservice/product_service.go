package productservice

import (
	"boiler-plate/internal/services"
	"boiler-plate/internal/storage"
	"boiler-plate/models"
)

type ProductService struct {
	services.BaseService[models.Product, models.ProductRequestDTO, models.ProductResponsDTO]
}

func NewProductService(storage storage.IStorage) services.IProductService {
	productService := &ProductService{}
	productService.NewBaseService(storage)
	return productService
}
