package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	productManager := NewProductManager()
	cart := NewBlankCart()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/products", func(context echo.Context) error {

		products := productManager.GetAll()
		return context.JSON(http.StatusOK, products)
	})

	e.POST("/add", func(context echo.Context) error {

		req := new(BuyingRequest)
		if err := context.Bind(req); err != nil {
			return nil
		}

		product := productManager.GetByName(req.name)

		cart.addToCart(product)

		return context.JSON(http.StatusOK, "Ok!")
	})



	e.Logger.Fatal(e.Start(":5000"))
}
