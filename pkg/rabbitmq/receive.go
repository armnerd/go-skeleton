package rabbitmq

import (
	"fmt"
	"time"

	"github.com/armnerd/go-skeleton/config"

	"github.com/streadway/amqp"
)

func Receive(queue string, msgCount int) (res []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover:%v\n", err)
			return
		}
	}()
	// 获取参数
	host := config.RabbitMQConfig.Host
	port := config.RabbitMQConfig.Port
	user := config.RabbitMQConfig.User
	pass := config.RabbitMQConfig.Pass
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"exchange_direct", // name
		"direct",          // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,            // queue name
		queue,             // routing key
		"exchange_direct", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// 获取消息
	res = make([]string, 0, msgCount)
	endtime := time.Now().Unix() + 2
	go func() {
		for {
			// 取够消息
			if len(res) >= msgCount {
				break
			}
			// 超时
			if time.Now().Unix() > endtime {
				break
			}
			d, ok := <-msgs
			if !ok {
				break
			}
			d.Ack(false)
			res = append(res, string(d.Body))
		}
	}()
	for {
		// 取够消息
		if len(res) >= msgCount {
			panic("enough")
		}
		// 超时
		if time.Now().Unix() > endtime {
			panic("time out")
		}
	}
}
