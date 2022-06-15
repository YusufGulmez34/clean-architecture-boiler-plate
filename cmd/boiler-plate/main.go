package main

import (
	"boiler-plate/config"
	"boiler-plate/internal/api/echoapi"
	"boiler-plate/internal/middlewares"
	"boiler-plate/internal/services/authservice"
	"boiler-plate/internal/services/productservice"
	"boiler-plate/pkg/validators"

	_ "boiler-plate/docs"

	"boiler-plate/internal/storage/database"
	"boiler-plate/pkg/db/mysql"
	"boiler-plate/pkg/global"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

// @title Boiler-Plate Backend Api
// @version 1.0
// @description Boiler-Plate backend api documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	config := config.Configuration{}
	config.ReadConfigFile()

	global.TokenExpiryTime = config.SystemConfig.TokenExpiryTime
	global.TokenSecretKey = config.SystemConfig.TokenSecretKey

	gormDb := mysql.Connect(config.DbConfig)
	store := database.NewGormDatabase(gormDb)

	e := echo.New()
	e.Validator = validators.NewCustomValidator(validator.New())
	e.GET("swagger/*", echoSwagger.WrapHandler)

	middlewareManager := middlewares.NewMiddlewareManager(store)

	//Auth
	authService := authservice.NewAuthService(store)
	authApi := echoapi.NewAuthApi(authService)
	authApi.Register(e.Group("auth"))

	//Product
	productService := productservice.NewProductService(store)
	productApi := echoapi.NewEchoProductApi(productService)
	productApi.Register(e.Group("products", middlewareManager.AuthControl))

	e.Start(":5000")
}
