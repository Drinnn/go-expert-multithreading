package main

func main() {
	ch := make(chan int)
	go publish(ch)
	read(ch)
}

func read(ch chan int) {
	for i := range ch {
		println(i)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
