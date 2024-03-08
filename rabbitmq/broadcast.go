package rabbitmq

import (
	"errors"

	"github.com/streadway/amqp"
)

// Broadcast 广播模式
func Broadcast(exchange, exchangeType string, content []byte) error {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return errors.New("链接rabbitmq失败")
	}

	channel, err := connection.Channel()
	if err != nil {
		return errors.New("创建管道失败")
	}

	if err = channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return errors.New("交换机声明失败")
	}

	if err = channel.Publish(
		exchange,     // publish to an exchange
		exchangeType, // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            content,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return errors.New("发送消息失败")
	}

	return nil
}

// BroadcastConsumption 广播消费
func BroadcastConsumption(exchange, exchangeType string, take Take) error {
	//链接rabbitmq
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return errors.New("链接rabbitmq失败")
	}

	//创建管道
	channel, err := connection.Channel()
	if err != nil {
		return errors.New("创建管道失败")
	}

	//一个交换机
	err = channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return errors.New("声明一个交换机失败")
	}

	//创建一个队列
	queue, err := channel.QueueDeclare(
		"",    // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return errors.New("创建一个交换机失败")
	}

	//绑定一个交换机
	err = channel.QueueBind(
		queue.Name,   // name of the queue
		exchangeType, // bindingKey
		exchange,     // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return errors.New("绑定一个交换机失败")
	}

	//消费
	deliveries, err := channel.Consume(
		queue.Name, // name
		"",         // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)

	for d := range deliveries {
		take(string(d.Body))
		//手动开启消费
		d.Ack(false)
	}
	return nil
}
