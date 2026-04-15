package main

import (
	"errors"
	"fmt"
)

func printOrderStatus(status string) {
	switch status {
	case "PENDING":
		fmt.Println("Don hang dang cho xu ly")
	case "PAID":
		fmt.Println("Don hang da thanh toan")
	case "CANCELLED":
		fmt.Println("Don hang da huy")
	default:
		fmt.Println("Trang thai khong hop le")
	}
}

func checkStock(skuCode string, stockMap map[string]int) {
	stock, ok := stockMap[skuCode]
	if !ok {
		fmt.Println("SKU khong ton tai trong he thong")
	} else {
		fmt.Println(stock)
	}
}

func processOrder(skuCode string, orderQty int, stockMap map[string]int) error {
	stock, ok := stockMap[skuCode]

	if !ok {
		return errors.New("SKU khong ton tai")
	} else if stock < orderQty {
		return errors.New("Khong du hang trong kho")
	} else {
		stockMap[skuCode] = stockMap[skuCode] - orderQty
	}
	return nil
}

func main() {
	skus := []int{1, 2, 3, 4, 5}
	stockData := make(map[string]int)
	stockData["SKU-001"] = 100
	stockData["SKU-002"] = 50
	stockData["SKU-003"] = 0

	totalPrice := 0
	for _, price := range skus {
		totalPrice += price
	}
	fmt.Println(totalPrice)

	checkStock("SKU-001", stockData)
	err := processOrder("SKU-001", 10, stockData)
	fmt.Println(err)
	fmt.Println(stockData["SKU-001"])
}
