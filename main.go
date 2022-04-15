package main

import (
	"fmt"
)

// Main function
func main() {

	var fibonacci func(n int) <-chan int
	// Define what number to seek
	fibTarget := [5]int{10, 15, 20, 25, 30}

	// Stores an anonymous function for finding Fibonacci numbers
	fibonacci = func(n int) <-chan int {
		result := make(chan int)

		// Goroutine by anonymous function
		go func() {
			// Reservation to close the channel
			defer close(result)
			// If less than 2, necessarily 1
			if n <= 2 {
				result <- 1
				return
			}
			// Find [n-1] by recursion
			f1 := <-fibonacci(n-1)
			// Find [n-2] by recursion
			f2 := <-fibonacci(n-2)
			// Add the [n-1] and [n-2] numbers to obtain the Fibonacci number
			result <- f1 + f2
		}()

		return result
	}

	// Find the Fibonacci numbers of the target
	for _, f := range fibTarget {
		fib := <-fibonacci(f)
		fmt.Printf("No.%d fibonacci is %d \n", f, fib)
	}
}

// =================================
//           Output Sample
// =================================
// ~ $ go build -o main main.go 
// ~ $ ./main 
// No.10 fibonacci is 55 
// No.15 fibonacci is 610 
// No.20 fibonacci is 6765 
// No.25 fibonacci is 75025 
// No.30 fibonacci is 832040 
