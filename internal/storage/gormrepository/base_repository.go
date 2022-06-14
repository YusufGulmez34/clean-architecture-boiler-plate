package gormrepository

import (
	"gorm.io/gorm"
)

type GormBaseRepository[TModel any] struct {
	db *gorm.DB
}

func (bp *GormBaseRepository[TModel]) NewGormBaseRepository(db *gorm.DB) {
	bp.db = db
}

func (bp *GormBaseRepository[TModel]) DB() *gorm.DB {
	return bp.db
}

func (bp *GormBaseRepository[TModel]) Create(model *TModel) error {
	return bp.db.Create(&model).Error
}

func (bp *GormBaseRepository[TModel]) FindAll() ([]*TModel, error) {
	var models []*TModel
	if err := bp.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (bp *GormBaseRepository[TModel]) FindByID(id uint) (*TModel, error) {
	var model TModel
	if err := bp.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (bp *GormBaseRepository[TModel]) Update(model *TModel) error {
	return bp.db.Save(&model).Error
}

func (bp *GormBaseRepository[TModel]) Delete(id uint) error {
	var model TModel
	return bp.db.Delete(&model, id).Error
}
