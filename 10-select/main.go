package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	rabbitMQFakeChannel := make(chan Message)
	kafkaFakeChannel := make(chan Message)

	msgIndex := int64(1)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			msg := Message{id: msgIndex, Msg: "Hello from Rabbit!"}
			rabbitMQFakeChannel <- msg
			atomic.AddInt64(&msgIndex, 1)
		}

	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			msg := Message{id: msgIndex, Msg: "Hello from Kafka!"}
			kafkaFakeChannel <- msg
			atomic.AddInt64(&msgIndex, 1)
		}
	}()

	for {
		select {
		case msg := <-rabbitMQFakeChannel:
			fmt.Println(fmt.Sprintf("Received from RabbitMQ: Id %d, Message: %s", msg.id, msg.Msg))
		case msg := <-kafkaFakeChannel:
			fmt.Println(fmt.Sprintf("Received from Kafka: Id %d, Message: %s", msg.id, msg.Msg))
		case <-time.After(3 * time.Second):
			println("timeout")
			//default:
			//	println("default")
		}
	}
}

type Message struct {
	id  int64
	Msg string
}
