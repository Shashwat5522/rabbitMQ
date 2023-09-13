package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)

	go func() {

		for d := range msgs {
			fmt.Printf("Received Messages:%s\n",d.Body)
		}
	}()

	fmt.Println("Successfully connected to our rabbitmq instance")
	fmt.Println("[*]-waiting for messages")
	<-forever
}
