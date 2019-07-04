package rabbitmq

import (
	log "alpha/config"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

// RabbitMQ stores rabbitmq's connection information
// it also handles disconnection (purpose of URL and QueueName storage)
type RabbitMQ struct {
	URL        string
	Exchange   string
	Conn       *amqp.Connection
	Chann      *amqp.Channel
	Queue      amqp.Queue
	closeChann chan *amqp.Error
	quitChann  chan bool
}

func (rmq *RabbitMQ) Load() error {
	var err error

	rmq.Conn, err = amqp.Dial(rmq.URL)
	if err != nil {
		log.Logger.Error("rabbit init",
			zap.Error(err),
		)
		return err
	}

	rmq.Chann, err = rmq.Conn.Channel()
	if err != nil {
		log.Logger.Error("rabbit init",
			zap.Error(err),
		)
		return err
	}

	rmq.closeChann = make(chan *amqp.Error)
	rmq.Conn.NotifyClose(rmq.closeChann)

	// declare exchange if not exist
	//err = rmq.Chann.ExchangeDeclare(rmq.Exchange, "direct", true, false, false, false, nil)
	//if err != nil {
	//	return errors.Wrapf(err, "declaring exchange %q", rmq.Exchange)
	//}
	exchange := strings.ToLower(rmq.Exchange)
	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	err = rmq.Chann.ExchangeDeclare(exchange, "x-delayed-message", true, false, false, false, args)
	if err != nil {
		log.Logger.Error("rabbit init",
			zap.Error(err),
		)
		return err
	}

	//err = declareConsumer(rmq)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (rmq *RabbitMQ) DeclareAndConsumer(queueName, routeKey string) (ds <-chan amqp.Delivery, err error) {
	queueName = strings.ToLower(queueName)
	routeKey = strings.ToLower(routeKey)
	delayedQueue, err := rmq.Chann.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Logger.Error("rabbit DeclareAndConsumer",
			zap.Error(err),
		)
		return nil, err
	}
	err = rmq.Chann.QueueBind(delayedQueue.Name, routeKey, rmq.Exchange, false, nil)
	if err != nil {
		log.Logger.Error("rabbit DeclareAndConsumer",
			zap.Error(err),
		)
		return nil, err
	}

	// Set our quality of service.  Since we're sharing 3 consumers on the same
	// channel, we want at least 2 messages in flight.
	err = rmq.Chann.Qos(2, 0, false)
	if err != nil {
		return nil, err
	}

	published, err := rmq.Chann.Consume(
		delayedQueue.Name,
		delayedQueue.Name+"_consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Logger.Error("rabbit DeclareAndConsumer",
			zap.Error(err),
		)
		return nil, err
	}
	return published, err
}

// declareConsumer declares all queues and bindings for the consumer
func declareConsumer(rmq *RabbitMQ) error {
	var err error

	// rmq.Queue, err = rmq.Chann.QueueDeclare("user-created-queue", true, false, false, false, nil)
	// if err != nil {
	// 	return err
	// }
	// err = rmq.Chann.QueueBind(rmq.Queue.Name, "user.event.create", rmq.Exchange, false, nil)
	// if err != nil {
	// 	return err
	// }

	delayedQueue, err := rmq.Chann.QueueDeclare("user-published-queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = rmq.Chann.QueueBind(delayedQueue.Name, "user.event.publish", "delayed", false, nil)
	if err != nil {
		return err
	}

	// Set our quality of service.  Since we're sharing 3 consumers on the same
	// channel, we want at least 2 messages in flight.
	err = rmq.Chann.Qos(2, 0, false)
	if err != nil {
		return err
	}

	published, err := rmq.Chann.Consume(
		"user-published-queue",
		"user-published-consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	go consume(published)

	return nil
}

func consume(ds <-chan amqp.Delivery) {
	for {

		select {
		case d, ok := <-ds:
			if !ok {
				return
			}
			d.Ack(false)
		}
	}
}

// Shutdown closes rabbitmq's connection
func (rmq *RabbitMQ) Shutdown() {
	rmq.quitChann <- true

	<-rmq.quitChann
}

func (rmq *RabbitMQ) Destroy() {
	rmq.Chann.Close()
	rmq.Conn.Close()
}

// handleDisconnect handle a disconnection trying to reconnect every 5 seconds
func (rmq *RabbitMQ) handleDisconnect() {
	for {
		select {
		case errChann := <-rmq.closeChann:
			if errChann != nil {
			}
		case <-rmq.quitChann:
			rmq.Conn.Close()
			rmq.quitChann <- true
			return
		}

		time.Sleep(5 * time.Second)

		if err := rmq.Load(); err != nil {
		}
	}
}

// Publish sends the given body on the routingKey to the channel
func (rmq *RabbitMQ) Publish(routingKey string, body []byte) error {
	return rmq.publish(rmq.Exchange, routingKey, body, 0)
}

// PublishWithDelay sends the given body on the routingKey to the channel with a delay
func (rmq *RabbitMQ) PublishWithDelay(routingKey string, body []byte, delay int64) error {
	return rmq.publish(rmq.Exchange, routingKey, body, delay)
}

func (rmq *RabbitMQ) publish(exchange string, routingKey string, body []byte, delay int64) error {
	headers := make(amqp.Table)
	routingKey = strings.ToLower(routingKey)
	exchange = strings.ToLower(exchange)
	if delay != 0 {
		headers["x-delay"] = delay * 1000
	}
	return rmq.Chann.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         body,
		Headers:      headers,
	})
}
