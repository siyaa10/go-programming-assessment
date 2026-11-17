package main

import (
	"errors"
	"fmt"
)

// Custom error structure
type ValidationError struct {
	Field  string
	Reason string
}

// Implement the built-in error interface
func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Reason
}

// Validate the age
func ValidateAge(age int) error {

	if age < 0 {
		return &ValidationError{
			Field:  "Age",
			Reason: "Age cannot be negative",
		}
	}

	if age > 120 {
		return &ValidationError{
			Field:  "Age",
			Reason: "Age cannot be greater than 120",
		}
	}

	return nil
}

// Demonstrate custom error checking

func runValidationDemo() {

	fmt.Println("\n===== AGE VALIDATION =====")

	var age int

	fmt.Print("Enter your age: ")
	_, scanErr := fmt.Scan(&age)

	if scanErr != nil {
		fmt.Println("Invalid input! Please enter a number.")
		return
	}

	err := ValidateAge(age)

	if err != nil {

		var validationErr *ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("Validation Error:", validationErr.Field)
			fmt.Println("Reason:", validationErr.Reason)
		}

	} else {
		fmt.Println("Age is valid!")
	}
}
