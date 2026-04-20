package main

import "fmt"

type Items interface {
	GetLineTotal() int
}

type SKU struct {
	Code     string
	Name     string
	Price    int
	StockQty int
}

func (s SKU) GetLineTotal() int {
	return s.Price
}

type OrderItem struct {
	SKUCode  string
	Price    int
	Quantity int
}

func (o OrderItem) GetLineTotal() int {
	return o.Price * o.Quantity
}

func PrintItemTotal(i Items) {
	fmt.Println("Total: ", i.GetLineTotal())
}

func main() {
	sku := OrderItem{
		SKUCode:  "SKU-001",
		Price:    100000,
		Quantity: 10,
	}

	PrintItemTotal(sku)
}