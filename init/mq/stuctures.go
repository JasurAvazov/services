package mq

import amqp "github.com/rabbitmq/amqp091-go"

type data struct {
	Id string `json:"id"`
}

type Conn struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

type response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

var (
	notFound = 404
	success = 200
)

const url  = "http://localhost:7077"