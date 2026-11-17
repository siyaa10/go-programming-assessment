package main

import (
	"fmt"
	"sync"
	"time"
)

func runGoroutineDemo() {

	fmt.Println("\n===== GOROUTINES =====")

	var count int

	fmt.Print("How many URLs do you want to download? ")
	fmt.Scan(&count)

	if count <= 0 {
		fmt.Println("Please enter a valid number.")
		return
	}

	urls := make([]string, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter URL %d: ", i+1)
		fmt.Scan(&urls[i])
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			fmt.Println("\nDownloading:", u)
			time.Sleep(2 * time.Second)
			fmt.Println("Download completed:", u)

		}(url)
	}

	wg.Wait()

	fmt.Println("\nAll downloads completed!")
}
