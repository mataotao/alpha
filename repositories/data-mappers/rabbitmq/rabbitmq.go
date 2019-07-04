package rabbitmq

import (
	"fmt"
	"strings"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Host    string
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (mq Rabbitmq) MarshalDefaults(v interface{}) {
	if rb, ok := v.(*Rabbitmq); ok {
		if rb.Host == "" {
			rb.Host = "amqp://guest:guest@localhost:5672/"
		}
	}
}

func (mq *Rabbitmq) Init() {
	if mq.conn == nil && mq.channel == nil {
		mq.conn, mq.channel = mq.NewChannel()
	}
}

func (mq *Rabbitmq) NewChannel() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(mq.Host)
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return conn, channel
}

func (mq *Rabbitmq) Destroy() (err error) {
	err = mq.channel.Close()
	err = mq.conn.Close()
	return
}

func (mq *Rabbitmq) Delay(routeKey string, body []byte, expire int64) error {
	err := mq.channel.Publish(
		"",
		strings.ToLower(routeKey),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
			Expiration:   fmt.Sprintf("%v", expire), // 延迟时间，实际上是过期时间
		},
	)
	if err != nil {
		return err
	}
	return err
}

func (mq *Rabbitmq) Consumer(delayQueue string) (err error, msgs <-chan amqp.Delivery) {
	delayQueue = strings.ToLower(delayQueue)
	// 声明主队列
	mainQueue, err := mq.channel.QueueDeclare(
		delayQueue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return
	}
	// 声明队列，这个队列不做消费,而是让消息变成死信后再进行转发，达到延迟队列的目的
	delayExchange := delayQueue + "_delay_exchange"
	decDelayQueue := delayQueue + "_delay"
	_, err = mq.channel.QueueDeclare(
		decDelayQueue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		amqp.Table{
			"x-dead-letter-exchange": delayExchange,
		},
	)

	if err != nil {
		return
	}
	// 声明exchange，接收延时队列消息
	err = mq.channel.ExchangeDeclare(
		delayExchange, // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return
	}
	// 将主监听队列和 exchange 绑定
	err = mq.channel.QueueBind(
		mainQueue.Name, // queue name
		"",             // routing key
		delayExchange,  // exchange
		false,
		nil)
	if err != nil {
		return
	}

	// 为了保证公平分发，不至于其中某个consumer一直处理，而其他不处理
	err = mq.channel.Qos(
		3,     // prefetchCount  在server收到consumer的ACK之前，预取的数量。为1，表示在没收到consumer的ACK之前，只会为其分发一个消息
		0,     // prefetchSize 大于0时，表示在收到consumer确认消息之前，将size个字节保留在网络中
		false, // global  true:Qos对同一个connection的所有channel有效； false:Qos对同一个channel上的所有consumer有效
	)
	if err != nil {
		return
	}

	// 消费
	msgs, err = mq.channel.Consume(
		mainQueue.Name, // queue
		"",             // consumer
		false,          // auto-ack   不进行自动ACK
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		return
	}
	return
}
