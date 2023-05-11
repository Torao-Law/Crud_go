package routes

import (
	"Product/handlers"
	"Product/pkg/mysql"
	"Product/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProduct)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", h.CreateProduct)
	e.PATCH("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)
}
