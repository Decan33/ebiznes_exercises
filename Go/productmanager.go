package main

import (
	"math/rand"

	"github.com/labstack/echo"
)

type CreateNewProduct struct {
	name  string `json:"name"`
	price int    `json:"price"`
}

type ChangeContentsProduct struct {
	name  string `json:"name"`
	price int    `json:"price"`
}

func (manager *ProductManager) GetAll() []Product {
	return manager.products
}

func (manager *ProductManager) Create(createNewProduct CreateNewProduct) Product {
	newProduct := Product{
		productId: rand.Intn(999),
		name:      createNewProduct.name,
		price:     createNewProduct.price,
	}

	manager.products = append(manager.products, newProduct)

	return newProduct
}

func (manager *ProductManager) ChangeContents(changeContentsProduct ChangeContentsProduct) Product {
	var index int = -1

	for i, element := range manager.products {
		if element.name == changeContentsProduct.name {
			index = i
		}
	}

	if index == -1 {
		return Product{}
	}

	manager.products[index].name = changeContentsProduct.name
	manager.products[index].price = changeContentsProduct.price

	return manager.products[index]
}

func (manager *ProductManager) Delete(productId int) error {
	index := -1

	for i, t := range manager.products {
		if t.productId == productId {
			index = i
			break
		}
	}

	if index == -1 {
		return echo.ErrNotFound
	}

	manager.products = append(manager.products[:index], manager.products[index+1:]...)

	return nil
}
