package mq

import amqp "github.com/rabbitmq/amqp091-go"

type Record interface {
	TakeId()
	LogToMQ(logLevel string, statusCode int)
	post(id string,i amqp.Delivery) string
}