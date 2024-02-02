package main

import (
	"fmt"
	"sync"
	"time"
)

func Task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		waitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)

	go Task("A", &waitGroup)
	go Task("B", &waitGroup)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d - Task C is running anonymously\n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
