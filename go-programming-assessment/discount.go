package main

import "fmt"

// Calculate discounted price
func CalculateDiscount(price float64, percentage float64) (float64, error) {

	if price < 0 {
		return 0, fmt.Errorf("price cannot be negative")
	}

	if percentage < 0 || percentage > 100 {
		return 0, fmt.Errorf("discount percentage must be between 0 and 100")
	}

	discountedPrice := price - (price * percentage / 100)

	return discountedPrice, nil
}

// Interactive discount calculator
func runDiscountDemo() {

	fmt.Println("\n===== DISCOUNT CALCULATOR =====")

	var price, percentage float64

	fmt.Print("Enter original price: ")
	fmt.Scan(&price)

	fmt.Print("Enter discount percentage: ")
	fmt.Scan(&percentage)

	result, err := CalculateDiscount(price, percentage)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Discounted Price: %.2f\n", result)
	}
}
