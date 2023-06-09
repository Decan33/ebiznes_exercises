package main

type Product struct {
	productId int
	name      string
	description string
	price     int
}

type BuyingRequest struct {
	name      string
	price     int
}

type CartItem struct {
	amount int
	product Product
}

type ProductManager struct {
	products []Product
}

type CartManager struct {
	items []CartItem
}

func NewProductManager() ProductManager {
	return ProductManager{
		products: []Product{
			{
				productId: 1,
				name: "Zajebista koszulka",
				description: "Nie ma co, musisz ja kupic",
				price: 999,
			},
			{
				productId: 2,
				name: "Zajebisty zegarek",
				description: "Pokazuje czas lepiej niz kazdy inny",
				price: 69,
			},
			{
				productId: 3,
				name: "Zaliczenie z Ebiznesu",
				description: "Fajnie byloby miec co nie",
				price: 420,
			},
			{
				productId: 4,
				name: "3.0 z PRiRu",
				description: "Nieosiągalne",
				price: 0,
			},
		},
	}
}

func NewBlankCart() CartManager {
	return CartManager{
		items: make([]CartItem, 0),
	}
}
