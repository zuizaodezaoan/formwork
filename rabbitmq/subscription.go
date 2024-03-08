package rabbitmq

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

// Product 简单订阅模式
func Product(queueName string, msgContent []byte) error {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return errors.New("链接rabbitmq失败")
	}

	channel, err := connection.Channel()
	if err != nil {
		return errors.New("创建管道失败")
	}

	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errors.New("创建一个队列失败")
	}

	if err = channel.Publish(
		"",         // publish to an exchange
		queue.Name, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            msgContent,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return errors.New("往队列发送消息失败")
	}

	return nil
}

type Take func(s string)

// Consumer 普通消费
func Consumer(queueName string, take Take) error {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return errors.New("链接rabbitmq失败")
	}

	channel, err := connection.Channel()
	if err != nil {
		return errors.New("创建管道失败")
	}

	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errors.New("创建一个队列失败")
	}

	deliveries, err := channel.Consume(
		queue.Name, // name
		"",         // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return errors.New("获取消息队列失败")
	}

	for d := range deliveries {
		//打印消息
		fmt.Println(string(d.Body))
		//消费
		take(string(d.Body))
		//手动改开启消费
		d.Ack(false)
	}

	return nil
}
