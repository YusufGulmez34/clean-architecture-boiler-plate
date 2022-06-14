package database

import (
	"boiler-plate/internal/storage"
	"boiler-plate/internal/storage/gormrepository/gormproductrepo"
	"boiler-plate/internal/storage/gormrepository/gormtokenrepo"
	"boiler-plate/internal/storage/gormrepository/gormuserrepo"

	"gorm.io/gorm"
)

type GormDatabase struct {
	db           *gorm.DB
	products     storage.IProductRepository
	tokenDetails storage.ITokenRepository
	users        storage.IUserRepository
}

func NewGormDatabase(db *gorm.DB) storage.IStorage {
	return &GormDatabase{
		db:           db,
		products:     gormproductrepo.NewGormProductRepository(db),
		tokenDetails: gormtokenrepo.NewGormTokenRepository(db),
		users:        gormuserrepo.NewGormUserRepository(db),
	}
}

func (gdb *GormDatabase) Products() storage.IProductRepository {
	return gdb.products
}

func (gdb *GormDatabase) TokenDetails() storage.ITokenRepository {
	return gdb.tokenDetails
}

func (gdb *GormDatabase) Users() storage.IUserRepository {
	return gdb.users
}
