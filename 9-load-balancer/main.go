package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	workersAmount := 100

	for i := 0; i < workersAmount; i++ {
		go initWorker(i, ch)
	}

	for i := 0; i < 1000; i++ {
		ch <- i
	}

}

func initWorker(workerId int, dataChannel <-chan int) {
	for x := range dataChannel {
		fmt.Println("Worker", workerId, "received", x)
		time.Sleep(time.Second)
	}
}
