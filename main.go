package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main(){
	fmt.Println("Go Rabbit ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Successfully connected to our rabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)

	err=ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte("Hello world"),
		},

	)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Sucessfully published message to queue")


}
