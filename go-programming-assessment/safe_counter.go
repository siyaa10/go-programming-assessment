package main

import (
	"fmt"
	"sync"
)

// SafeCounter structure
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment the counter safely
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Get the current counter value safely
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

// Demonstrate concurrent counter

func runSafeCounterDemo() {
	fmt.Println("\n===== SAFE COUNTER =====")

	var people int
	var times int

	fmt.Print("Enter number of goroutines: ")
	_, err := fmt.Scan(&people)
	if err != nil || people <= 0 {
		fmt.Println("Please enter a valid positive number.")
		return
	}

	fmt.Print("Enter increments per goroutine: ")
	_, err = fmt.Scan(&times)
	if err != nil || times <= 0 {
		fmt.Println("Please enter a valid positive number.")
		return
	}

	var counter SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < people; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < times; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println("\nProcessing completed!")
	fmt.Println("Number of Goroutines:", people)
	fmt.Println("Increments per Goroutine:", times)
	fmt.Println("Expected Counter Value:", people*times)
	fmt.Println("Final Counter Value:", counter.Value())
}
