package mq

import "encoding/json"

func (c *Conn) TakeId(){
	q, err := c.Ch.QueueDeclare(
		"queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	Error(err, "Failed to declare a queue")

	msgs, err := c.Ch.Consume(
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
		c.post(d.Id, i)
	}
}
