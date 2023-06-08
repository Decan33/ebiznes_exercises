package main

type Product struct {
	productId int
	name      string
	price     int
}

type ProductManager struct {
	products []Product
}

func NewProductManager() ProductManager {
	return ProductManager{
		products: make([]Product, 0),
	}
}
