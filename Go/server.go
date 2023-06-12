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
	registeredUsersDatabase := NewDatabase()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/products", func(context echo.Context) error {

		products := productManager.GetAll()
		fmt.Println(products)
		return context.JSON(http.StatusOK, products)
	})

	e.GET("/cart", func(context echo.Context) error {

		items := cart.GetAllCartItems()
		return context.JSON(http.StatusOK, items)
	})

	e.POST("/add", func(context echo.Context) error {

		req := new(BuyingRequest)
		if err := context.Bind(req); err != nil {
			return err
		}

		product, secondErr := productManager.GetByName(req.Name)

		if secondErr != nil {
			return context.JSON(http.StatusNotFound, "asdf")
		}

		cart.AddToCart(product)

		return context.JSON(http.StatusOK, "Ok!")
	})

	e.POST("/pay", func(ctx echo.Context) error {

		req := new(PaymentRequest)

		if err := ctx.Bind(req); err != nil {
			return nil
		}

		paymentAmount := req.Amount

		items := cart.items
		sum := 0
		for _, item := range items {
			sum += item.Amount * item.Product.Price
		}

		if paymentAmount < sum {
			return ctx.JSON(http.StatusBadRequest, "Not enough money")
		}

		remainder := paymentAmount - sum
		cart.items = nil
		return ctx.JSON(http.StatusOK, fmt.Sprintf("We have to return you %d zlotys!", remainder))
	})

	e.POST("/log", func(ctx echo.Context) error {

		req := new(LogInRequest)

		if err := ctx.Bind(req); err != nil {
			return nil
		}

		login := req.Login
		password := req.Password

		credentials := LogInCredentials{
			login:    login,
			password: password,
		}

		if registeredUsersDatabase.LogUser(credentials) {
			return ctx.JSON(http.StatusOK, "User logged in successfully!")
		}
		return ctx.JSON(http.StatusNotFound, "We could not find the user :(")
	})

	e.Logger.Fatal(e.Start(":5000"))
}
