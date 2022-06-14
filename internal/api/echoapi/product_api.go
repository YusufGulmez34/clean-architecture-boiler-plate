package echoapi

import (
	"boiler-plate/internal/services"
	"boiler-plate/models"
	"boiler-plate/pkg/constants"
	"boiler-plate/pkg/result"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoProductApi struct {
	productService services.IProductService
}

func NewEchoProductApi(productService services.IProductService) *EchoProductApi {
	return &EchoProductApi{productService: productService}
}

func (p *EchoProductApi) Register(g *echo.Group) {
	g.POST("", p.Add)
	g.GET("", p.GetAll)
	g.GET("/:id", p.GetByID)
	g.PUT("/:id", p.Update)
	g.DELETE("/:id", p.Delete)
}

// @Tags ProductApi
// @Security ApiKeyAuth
// @Summary Get Product List
// @Success 200 string Success "{"success":true,"data":"","msg":"Success"}"
// @Router /products [get]
func (p *EchoProductApi) GetAll(c echo.Context) error {
	productListDTO, err := p.productService.GetAll()
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}
	return result.NewDataResult(c, http.StatusOK, constants.Success, productListDTO)
}

// @Tags ProductApi
// @Security ApiKeyAuth
// @Summary Get Product with ID
// @Param id path string true "Product ID"
// @Success 200 string Success "{"success":true,"data":"","msg":"Success"}"
// @Router /products/{id} [get]
func (p *EchoProductApi) GetByID(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	productDTO, err := p.productService.GetByID(uint(productID))
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewDataResult(c, http.StatusOK, constants.Success, productDTO)
}

// @Tags ProductApi
// @Security ApiKeyAuth
// @Summary Create Product
// @Param data body models.ProductRequestDTO true "Product Request DTO"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /products [post]
func (p *EchoProductApi) Add(c echo.Context) error {
	var productDTO models.ProductRequestDTO
	if err := c.Bind(&productDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	if err := p.productService.Add(productDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewSuccessResult(c, http.StatusCreated, constants.Success)
}

// @Tags ProductApi
// @Security ApiKeyAuth
// @Summary Delete Product
// @Param id path string true "Product ID"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /products/{id} [delete]
func (p *EchoProductApi) Delete(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	if err = p.productService.Delete(uint(productID)); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewSuccessResult(c, http.StatusOK, constants.Success)
}

// @Tags ProductApi
// @Security ApiKeyAuth
// @Summary Update Product
// @Param id path string true "Product ID"
// @Param data body models.ProductRequestDTO true "Product Request DTO"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /products/{id} [put]
func (p *EchoProductApi) Update(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	var productDTO models.ProductRequestDTO
	if err := c.Bind(&productDTO); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ParameterError, err)
	}

	if err = p.productService.Update(productDTO, uint(productID)); err != nil {
		return result.NewErrorResult(c, http.StatusBadRequest, constants.ServiceError, err)
	}

	return result.NewSuccessResult(c, http.StatusOK, constants.Success)

}
