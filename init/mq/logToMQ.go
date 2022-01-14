package mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
)

func LogToMQ(logLevel string, statusCode int) {
	fmt.Println(logLevel,"<--- loglevel\n",statusCode,"<---- status code")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Error(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Error(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	Error(err, "Failed to declare an exchange")

	body := strconv.Itoa(statusCode)
	fmt.Println("body ---->",body)
	err = ch.Publish(
		"logs_direct",          // exchange
		logLevel,                       // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	Error(err, "Failed to publish a message")
}

func Error(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}