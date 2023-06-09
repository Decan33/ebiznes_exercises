package main

import "golang.org/x/exp/slices"

func (manager *ProductManager) GetAll() []Product {
	return manager.products
}

func (manager *ProductManager) GetByName(name string) Product {
	index := slices.IndexFunc(manager.products, func(product Product) bool {
		return product.name == name
	})

	return manager.products[index]
}

func (cart *CartManager) addToCart(product Product) {
	if len(cart.items) == 0 {
		cart.items = append(cart.items, CartItem{
			amount: 1,
			product: product,
		})
	}

	index := slices.IndexFunc(cart.items, func(item CartItem) bool {
		return item.product == product
	})

	if index == -1 {
		cart.items = append(cart.items, CartItem{
			amount: 1,
			product: product,
		})
	}

	cart.items[index].amount += 1
}
