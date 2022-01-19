package mq

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Conn) post(id string,i amqp.Delivery) string {
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
		c.LogToMQ("error", notFound)
		i.Nack(false,false)
	case resp.ErrorCode == success:
		fmt.Println(resp.Data)
		fmt.Println("success")
		c.LogToMQ("info", success)
		i.Ack(false)
	default:
		i.Nack(false,true)
		fmt.Println("internal error")
	}
	return strconv.Itoa(resp.ErrorCode)
}
