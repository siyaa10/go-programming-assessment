package main

import (
	"math"
	"testing"
)

func TestCalculateDiscount(t *testing.T) {

	tests := []struct {
		name       string
		price      float64
		percentage float64
		expected   float64
		wantError  bool
	}{
		{"Valid discount", 1000, 20, 800, false},
		{"Full discount", 500, 100, 0, false},
		{"Negative price", -100, 20, 0, true},
		{"Negative percentage", 1000, -10, 0, true},
		{"Percentage above 100", 1000, 150, 0, true},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result, err := CalculateDiscount(
				test.price,
				test.percentage,
			)

			if (err != nil) != test.wantError {
				t.Errorf("Expected error: %v, got: %v",
					test.wantError, err)
				return
			}

			if err == nil && math.Abs(result-test.expected) > 0.000001 {
				t.Errorf("Expected %.2f, got %.2f",
					test.expected, result)
			}
		})
	}
}
