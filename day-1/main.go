package main

import (
	"errors"
	"fmt"
)

func calculateLineTotal(price int, quantity int) (int, error) {

	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than 0")
	}
	return price * quantity, nil
}

func countValidSKUs(skus [3]string) int {
	count := 0
	for i := 0; i < len(skus); i++ {
		if len(skus[i]) > 3 {
			count += 1
		}
	}
	return count
}

func main() {
	productName := "Keyboard"
	brand := "abv"
	isActive := true

	fmt.Println(productName)
	fmt.Println(brand)
	fmt.Println(isActive)

	total, err := calculateLineTotal(1000000, 2)
	// Vì chỗ này err đang được định nghĩa là: nếu không có lỗi nó sẽ là nil, còn có lỗi nó sẽ khác nil
	// suy ra là err != nil có nghĩa là đang có lỗi
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(total)
	}

	skus := [3]string{
		"IP-15-PRO",
		"SS",
		"MAC-M3-16GB",
	}

	for _, sku := range skus {
		fmt.Println(sku)
	}

	fmt.Println("Tong so skus co len > 3:", countValidSKUs(skus))
}
