package services

import (
	"boiler-plate/internal/storage"
	"boiler-plate/pkg/helpers"
	"errors"
	"reflect"

	"github.com/jinzhu/copier"
)

type BaseService[TModel any, TRequest any, TResponse any] struct {
	storage storage.IStorage
}

func (b *BaseService[TModel, TRequest, TResponse]) NewBaseService(storage storage.IStorage) {
	b.storage = storage
}

func (b *BaseService[TModel, TRequest, TResponse]) Storage() storage.IStorage {
	return b.storage
}

func (b *BaseService[TModel, TRequest, TResponse]) Add(modelDTO TRequest) error {
	var addedModel TModel
	if err := copier.Copy(&addedModel, modelDTO); err != nil {
		return err
	}

	vStorage := reflect.ValueOf(b.storage)

	vAddedModel := reflect.ValueOf(&addedModel)
	name := reflect.TypeOf(addedModel).Name()

	var params []reflect.Value
	params = append(params, vAddedModel)

	methodName := helpers.MethodName(name)
	modelRepo := vStorage.MethodByName(methodName)

	repo := modelRepo.Call(nil)
	err := repo[0].MethodByName("Create").Call(params)[0].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
}

func (b *BaseService[TModel, TRequest, TResponse]) GetAll() ([]*TResponse, error) {
	var model TModel
	vStorage := reflect.ValueOf(b.storage)
	name := reflect.TypeOf(model).Name()
	methodName := helpers.MethodName(name)
	modelRepo := vStorage.MethodByName(methodName)

	repo := modelRepo.Call(nil)
	result := repo[0].MethodByName("FindAll").Call(nil)
	modelList := result[0].Interface()
	err := result[1].Interface()
	if err != nil {
		return nil, err.(error)
	}
	var modelListDTO []*TResponse
	if err := copier.Copy(&modelListDTO, modelList); err != nil {
		return nil, err
	}
	if modelListDTO == nil {
		return nil, errors.New("List is empty")
	}
	return modelListDTO, nil
}

func (b *BaseService[TModel, TRequest, TResponse]) GetByID(id uint) (*TResponse, error) {
	var model TModel
	vStorage := reflect.ValueOf(b.storage)
	name := reflect.TypeOf(model).Name()

	methodName := helpers.MethodName(name)
	modelRepo := vStorage.MethodByName(methodName)

	repo := modelRepo.Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(id))

	result := repo[0].MethodByName("FindByID").Call(params)
	models := result[0].Interface()
	err := result[1].Interface()
	if err != nil {
		return nil, err.(error)
	}
	var modelDTO TResponse
	if err := copier.Copy(&modelDTO, models); err != nil {
		return nil, err
	}
	return &modelDTO, nil
}

func (b *BaseService[TModel, TRequest, TResponse]) Delete(id uint) error {
	var model TModel
	vStorage := reflect.ValueOf(b.storage)
	name := reflect.TypeOf(model).Name()

	methodName := helpers.MethodName(name)
	modelRepo := vStorage.MethodByName(methodName)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(id))

	repo := modelRepo.Call(nil)
	result := repo[0].MethodByName("FindByID").Call(params)

	err := result[1].Interface()

	if err != nil {
		return err.(error)
	}

	err = repo[0].MethodByName("Delete").Call(params)[0].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
}

func (b *BaseService[TModel, TRequest, TResponse]) Update(modelDTO TRequest, id uint) error {
	var addedModel TModel
	if err := copier.Copy(&addedModel, modelDTO); err != nil {
		return err
	}

	vStorage := reflect.ValueOf(b.storage)

	name := reflect.TypeOf(addedModel).Name()

	var params []reflect.Value
	params = append(params, reflect.ValueOf(id))

	methodName := helpers.MethodName(name)
	modelRepo := vStorage.MethodByName(methodName)

	repo := modelRepo.Call(nil)
	result := repo[0].MethodByName("FindByID").Call(params)

	model := result[0].Interface().(*TModel)
	err := result[1].Interface()
	if err != nil {
		return err.(error)
	}

	if err := copier.Copy(&model, modelDTO); err != nil {
		return err
	}
	params[0] = reflect.ValueOf(model)
	err = repo[0].MethodByName("Update").Call(params)[0].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
}
