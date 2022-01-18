package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"third/mq"
)

const url  = "http://localhost:7077"

type data struct {
	Id string `json:"id"`
}

func takeMsgs(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Error(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Error(err, "Failed to open a channel")
	defer ch.Close()

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
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Error(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		d := data{}
		for i := range msgs{
			_ = json.Unmarshal(i.Body, &d)
			post(d.Id)
		}
	}()
	<-forever
}

type response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

func post(id string) {
	res,err :=http.Get(fmt.Sprintf("%s/record/%s",url,id))
	if err != nil{
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print("\n\n")
	resp := response{}
	_ = json.Unmarshal(all, &resp)
	switch {
	case resp.ErrorCode == notFound:
		fmt.Println("not found")
		mq.LogToMQ("error", notFound)
	case resp.ErrorCode == success:
		fmt.Println(resp.Data)
		fmt.Println("success")
		mq.LogToMQ("info", success)
	default:
		mq.SendToQueue(id)
		fmt.Println("internal error")
	}
}

func main() {
	takeMsgs()
}


// 1) дергаю первый сервис который мне отправит статус код и тело
// 2) проверяю статус код
// типы статусов 200 , 404 , 500
// 3) если 200 то мы выводим тело и отправляем статус код и тело во второй сервис
// 3) если 404 то мы выводим то что не нашло и логируем код в сервис два
// 3) если 500 мы будем занова отправлять в сервис(1) до того момента пока не найдет. Cохраняем тело и код в логгер

// сейчас нужно взять код и тело у сервиса один

// 1) функция post обеспечивает связь с первым сервисом
// 2,3) функция linkBetweenLogger обеспечивает отправку данных

type Res struct {
	Code string `json:"code,omitempty"`
	Body string `json:"body,omitempty"`
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