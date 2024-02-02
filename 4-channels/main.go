package main

import "fmt"

// Main thread
func main() {
	ch := make(chan string) // Empty

	// Thread 2
	go func() {
		ch <- "Hello, World!" // Fill channel
	}()

	msg := <-ch // Empty again

	fmt.Println(msg)
}
