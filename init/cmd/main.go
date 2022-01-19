package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"third/mq"
)

const url  = "http://localhost:7077"

type data struct {
	Id string `json:"id"`
}

func takeMsgs(ch *amqp.Channel){
	q, err := ch.QueueDeclare(
		"queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	Error(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Error(err, "Failed to register a consumer")


	d := data{}
	for i := range msgs {
		_ = json.Unmarshal(i.Body, &d)
		post(d.Id, i)
	}

}

type response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

func post(id string,i amqp.Delivery) string {
	res,err :=http.Get(fmt.Sprintf("%s/record/%s",url,id))
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	resp := response{}
	_ = json.Unmarshal(all, &resp)
	fmt.Print("\n\n")
	switch {
	case resp.ErrorCode == notFound:
		fmt.Println("not found")
		mq.LogToMQ("error", notFound)
		i.Nack(false,false)
	case resp.ErrorCode == success:
		fmt.Println(resp.Data)
		fmt.Println("success")
		mq.LogToMQ("info", success)
		i.Ack(false)
	default:
		i.Nack(false,true)
		fmt.Println("internal error")
	}
	return strconv.Itoa(resp.ErrorCode)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Error(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Error(err, "Failed to open a channel")
	defer ch.Close()
	takeMsgs(ch)
}

var (
	notFound = 404
	success = 200
)

func Error(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}