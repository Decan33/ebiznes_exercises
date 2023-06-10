package main

import (
	"fmt"
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

	e.GET("/cart", func(context echo.Context) error {

		items := cart.GetAllCartItems()
		return context.JSON(http.StatusOK, items)
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

	e.POST("/pay", func(ctx echo.Context) error {

		req := new(PaymentRequest)

		if err := ctx.Bind(req); err != nil {
			return nil
		}

		paymentAmount := req.amount

		items := cart.items
		sum := 0
		for _, item := range items {
			sum += item.amount * item.product.price
		}

		if paymentAmount < sum {
			return ctx.JSON(http.StatusBadRequest, "Not enough money")
		}

		remainder := paymentAmount - sum
		return ctx.JSON(http.StatusOK, fmt.Sprintf("We have to return you %d zlotys!", remainder))
	})

	e.Logger.Fatal(e.Start(":5000"))
}
