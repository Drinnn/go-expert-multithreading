package main

import "sync"

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go read(ch, &wg)

	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func read(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		println(i)
		wg.Done()
	}
}
