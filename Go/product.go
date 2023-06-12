package main

type Product struct {
	ProductId   int    `json:"productId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type BuyingRequest struct {
	Name string `json:"name"`
}

type PaymentRequest struct {
	Amount int `json:"amount"`
}

type LogInRequest struct {
	Login    string `json:"amount"`
	Password string `json:"password"`
}

type LogInCredentials struct {
	login    string
	password string
}

type CartItem struct {
	Amount  int     `json:"amount"`
	Product Product `json:"product"`
}

type ProductManager struct {
	products []Product
}

type CartManager struct {
	items []CartItem
}

type TotallyNotLoginDatabase struct {
	registeredUsers []LogInCredentials
}

func NewDatabase() TotallyNotLoginDatabase {
	return TotallyNotLoginDatabase{
		registeredUsers: []LogInCredentials{
			{
				login:    "admin",
				password: "admin",
			},
		},
	}
}

func NewProductManager() ProductManager {
	manager := ProductManager{
		products: []Product{
			{
				ProductId:   1,
				Name:        "Zajebista koszulka",
				Description: "Nie ma co, musisz ja kupic",
				Price:       999,
			},
			{
				ProductId:   2,
				Name:        "Zajebisty zegarek",
				Description: "Pokazuje czas lepiej niz kazdy inny",
				Price:       69,
			},
			{
				ProductId:   3,
				Name:        "Zaliczenie z Ebiznesu",
				Description: "Fajnie byloby miec co nie",
				Price:       420,
			},
			{
				ProductId:   4,
				Name:        "3.0 z PRiRu",
				Description: "Nieosiągalne",
				Price:       0,
			},
		},
	}
	return manager
}

func NewBlankCart() CartManager {
	return CartManager{
		items: make([]CartItem, 0),
	}
}
