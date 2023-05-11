package handlers

import (
	productsdto "Product/dto/product"
	dto "Product/dto/result"
	"Product/models"
	"Product/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerproduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerproduct {
	return &handlerproduct{ProductRepository}
}

func (h *handlerproduct) FindProduct(c echo.Context) error {
	films, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerproduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// film.ThumbnailFilm = path_file + film.ThumbnailFilm

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerproduct) CreateProduct(c echo.Context) error {
	request := new(productsdto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// userLogin := c.Get("userLogin")
	// categoryId := userLogin.(jwt.MapClaims)["id"].(float64)

	film := models.Product{
		Name: request.Name,
	}

	film, err = h.ProductRepository.CreateProduct(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film, _ = h.ProductRepository.GetProduct(film.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}
func (h *handlerproduct) UpdateProduct(c echo.Context) error {
	request := new(productsdto.UpdateProductRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		category.Name = request.Name
	}

	// if request.Email != "" {
	// 	user.Email = request.Email
	// }

	// if request.Password != "" {
	// 	user.Password = request.Password
	// }

	data, err := h.ProductRepository.UpdateProduct(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseproduct(data)})
}

func (h *handlerproduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseproduct(data)})
}

func convertResponseproduct(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		ID:   u.ID,
		Name: u.Name,
		// Email:    u.Email,
		// Password: u.Password,
	}
}
