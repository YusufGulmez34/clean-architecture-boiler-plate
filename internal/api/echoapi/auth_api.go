package echoapi

import (
	"boiler-plate/internal/services"
	"boiler-plate/models"
	"boiler-plate/pkg/constants"
	"boiler-plate/pkg/result"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthApi struct {
	authService services.IAuthService
}

func NewAuthApi(authService services.IAuthService) *AuthApi {
	return &AuthApi{authService: authService}
}

func (a *AuthApi) Register(g *echo.Group) {
	g.POST("/register", a.UserRegister)
	g.POST("/login", a.UserLogin)

}

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary User Register
// @Param data body models.UserRegisterDTO true "Register DTO"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/register [post]
func (a *AuthApi) UserRegister(c echo.Context) error {
	var userRegisterDTO models.UserRegisterDTO

	if err := c.Bind(&userRegisterDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	if err := c.Validate(&userRegisterDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ValidationError, err)
	}

	if err := a.authService.Register(userRegisterDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewSuccessResult(c, http.StatusCreated, constants.Success)
}

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary User Login
// @Param data body models.UserLoginDTO true "Register DTO"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/login [post]
func (a *AuthApi) UserLogin(c echo.Context) error {
	var userLoginDTO models.UserLoginDTO

	if err := c.Bind(&userLoginDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	if err := c.Validate(&userLoginDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ValidationError, err)
	}

	tokenString, err := a.authService.Login(userLoginDTO)
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewDataResult(c, http.StatusOK, constants.Success, echo.Map{"token": tokenString})

}
