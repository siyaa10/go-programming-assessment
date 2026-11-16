package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculateTotal(marks []int) int {
	total := 0
	for _, mark := range marks {
		total += mark
	}
	return total
}

func calculateGrade(average float64) string {
	if average >= 80 {
		return "A"
	} else if average >= 60 {
		return "B"
	} else if average >= 40 {
		return "C"
	} else {
		return "F"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var name string
	marks := make([]int, 3)
	var grade string
	var result string

	fmt.Print("Enter student name : ")
	scanner.Scan()
	name = scanner.Text()

	for i := 0; i < 3; i++ {

		for {
			fmt.Printf("Enter mark %d: ", i+1)

			if !scanner.Scan() {
				fmt.Println("Unable to read input.")
				return
			}
			input := scanner.Text()

			mark, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid input! Please enter a number.")
				continue
			}

			if mark < 0 || mark > 100 {
				fmt.Println("Invalid mark! Enter between 0 and 100.")
				continue
			}

			marks[i] = mark
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("\nStudent:", name)
	fmt.Println("Mark 1:", marks[0])
	fmt.Println("Mark 2:", marks[1])
	fmt.Println("Mark 3:", marks[2])
	total := CalculateTotal(marks)
	fmt.Println("\nTotal Marks : ", total)

	average := float64(total) / 3
	fmt.Printf("Average Marks : %.2f\n", average)

	grade = calculateGrade(average)

	if average >= 40 {
		result = "Pass"
	} else {
		result = "Fail"
	}

	fmt.Println("Grade : ", grade)
	fmt.Println("Result : ", result)
}
