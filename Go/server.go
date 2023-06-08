package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	productManager := NewProductManager()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/", func(context echo.Context) error {
		products := productManager.GetAll()
		return context.JSON(http.StatusOK, products)
	})

	e.POST("/create", func(ctx echo.Context) error {
		body := CreateNewProduct{}

		err := ctx.Bind(&body)
		if err != nil {
			return err
		}

		product := productManager.Create(body)
		return ctx.JSON(http.StatusCreated, product)
	})

	e.PUT("/change", func(ctx echo.Context) error {
		body := ChangeContentsProduct{}

		err := ctx.Bind(&body)
		if err != nil {
			return err
		}

		productManager.ChangeContents(body)

		return nil
	})

	e.DELETE("/delete/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		realId, realError := strconv.Atoi(id)

		if realError != nil {
			ctx.Error(realError)
			return realError
		}

		err := productManager.Delete(realId)

		if err != nil {
			ctx.Error(err)
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
