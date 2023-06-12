package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func (manager *ProductManager) GetAll() []Product {
	return manager.products
}

func (manager *ProductManager) GetByName(name string) (Product, error) {
	index := slices.IndexFunc(manager.products, func(product Product) bool {
		return product.Name == name
	})

	if index == -1 {
		return Product{}, fmt.Errorf("product not found")
	}

	return manager.products[index], nil
}

func (cart *CartManager) AddToCart(product Product) {
	if len(cart.items) == 0 {
		cart.items = append(cart.items, CartItem{
			Amount:  1,
			Product: product,
		})

		return
	}

	index := slices.IndexFunc(cart.items, func(item CartItem) bool {
		return item.Product == product
	})

	if index == -1 {
		cart.items = append(cart.items, CartItem{
			Amount:  1,
			Product: product,
		})
		return
	}

	cart.items[index].Amount += 1
}

func (cart *CartManager) GetAllCartItems() []CartItem {
	return cart.items
}

func (database *TotallyNotLoginDatabase) LogUser(credentials LogInCredentials) bool {
	return slices.Contains[LogInCredentials](database.registeredUsers, credentials)
}
