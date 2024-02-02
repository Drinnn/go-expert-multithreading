package main

func main() {
	ch := make(chan string)
	go write(ch, "Hello, World!")

	println(read(ch))
}

// set the channel to only receive data
func write(ch chan<- string, msg string) {
	ch <- msg
}

// set the channel to only send data
func read(ch <-chan string) string {
	return <-ch
}
