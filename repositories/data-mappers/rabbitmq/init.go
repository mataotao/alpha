package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"net"
	"strconv"
	"time"
)

var Url string

func Init() {
	Url = fmt.Sprintf("amqp://%s:%s@%s:%d/", viper.GetString("rabbitmq.user"), viper.GetString("rabbitmq.pwd"), viper.GetString("rabbitmq.addr"), viper.GetInt("rabbitmq.port"))

}

func GetConn(queue string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.DialConfig(Url, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 10*time.Second)
		},
	})
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	if _, err := ch.QueueDeclare(queue, false, false, false, false, nil); err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

func Consume(queue string) (*amqp.Connection, <-chan amqp.Delivery, error) {
	conn, ch, err := GetConn(queue)
	if err != nil {
		return nil, nil, err
	}

	if err := ch.Qos(1, 0, false); err != nil {
		return nil, nil, err
	}

	deliveries, err := ch.Consume(queue, "", false, true, false, false, nil)

	return conn, deliveries, err
}

func Producer(queue, body string) error {
	conn, pub, err := GetConn(queue)
	if err != nil {
		return err
	}
	defer conn.Close()
	defer pub.Close()
	if err != nil {
		return err
	}
	if err := pub.Publish("", queue, false, false, amqp.Publishing{
		Body: []byte(body),
	}); err != nil {
		return err
	}
	return nil
}

/**
延迟队列
*/
func ConsumeDelay(queue string) (*amqp.Connection, <-chan amqp.Delivery, error) {
	conn, err := amqp.Dial(Url)
	if err != nil {
		return conn, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return conn, nil, err
	}

	exQueue := fmt.Sprintf("%s_ex", queue)
	delayQueue := fmt.Sprintf("%s_delay", queue)

	err = ch.ExchangeDeclare(exQueue, "fanout", true, false, false, false, nil)
	if err != nil {
		return conn, nil, err
	}

	q, err := ch.QueueDeclare(queue, false, false, true, false, nil)
	if err != nil {
		return conn, nil, err
	}
	_, err = ch.QueueDeclare(delayQueue, false, false, true, false,
		amqp.Table{
			"x-dead-letter-exchange": exQueue,
		},
	)
	if err != nil {
		return conn, nil, err
	}
	err = ch.QueueBind(q.Name, "", exQueue, false, nil)
	if err != nil {
		return conn, nil, err
	}
	msg, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return conn, msg, err
	}
	return conn, msg, nil
}

func ProducerDelay(queue string, body []byte, t int) error {
	conn, err := amqp.Dial(Url)
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.Publish("", fmt.Sprintf("%s_delay", queue), false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			Expiration:  strconv.Itoa(t * 1000),
		})
	if err != nil {
		return err
	}
	return nil
}
