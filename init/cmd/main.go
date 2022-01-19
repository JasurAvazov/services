package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"third/mq"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	mq.Error(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	mq.Error(err, "Failed to open a channel")
	defer ch.Close()
	c := mq.Conn{Conn: conn, Ch: ch}
	c.TakeId()
}