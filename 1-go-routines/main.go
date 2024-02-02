package main

import (
	"fmt"
	"time"
)

func Task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go Task("A")
	go Task("B")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d - Task C is running anonymously\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
}
