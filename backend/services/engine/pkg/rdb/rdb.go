package rdb

import (
	"errors"

	redis "github.com/redis/go-redis/v9"
	// amqp "github.com/rabbitmq/amqp091-go"
)

// Redis is our wrapper logic to instrument redis calls
type Redis struct {
	addresses []string
	client    redis.UniversalClient
}

// NewClient is used to create new Redis type. This type
// encapsulates our interaction with redis and provides instrumentation with new relic.
func NewClient(addresses []string) (*Redis, error) {
	if len(addresses) == 0 {
		return nil, errors.New("redis addresses list cannot be empty")
	}

	for _, dsn := range addresses {
		if dsn == "" {
			return nil, errors.New("dsn cannot be empty")
		}
	}

	var client redis.UniversalClient

	if len(addresses) == 1 {
		opts, err := redis.ParseURL(addresses[0])
		if err != nil {
			return nil, err
		}

		client = redis.NewClient(opts)
	} else {
		client = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: addresses,
		})
	}

	return &Redis{addresses: addresses, client: client}, nil
}

// Client is to return underlying redis interface
func (r *Redis) Client() redis.UniversalClient {
	return r.client
}

// MakeRedisClient is used to fulfill asynq's interface
func (r *Redis) MakeRedisClient() interface{} {
	return r.client
}

// RabbitMQ:

// // RabbitMQ is our wrapper logic to instrument RabbitMQ calls
// type RabbitMQ struct {
// 	url     string
// 	conn    *amqp.Connection
// 	channel *amqp.Channel
// }

// // NewClient is used to create a new RabbitMQ type. This type
// // encapsulates our interaction with RabbitMQ and provides instrumentation.
// func NewClient(url string) (*RabbitMQ, error) {
// 	if url == "" {
// 		return nil, errors.New("RabbitMQ URL cannot be empty")
// 	}

// 	conn, err := amqp.Dial(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	channel, err := conn.Channel()
// 	if err != nil {
// 		conn.Close()
// 		return nil, err
// 	}

// 	return &RabbitMQ{url: url, conn: conn, channel: channel}, nil
// }

// // Channel is to return the underlying RabbitMQ channel
// func (r *RabbitMQ) Channel() *amqp.Channel {
// 	return r.channel
// }

// // MakeRabbitMQClient is used to fulfill an interface if needed
// func (r *RabbitMQ) MakeRabbitMQClient() interface{} {
// 	return r.channel
// }

// // Close is used to close the RabbitMQ connection and channel
// func (r *RabbitMQ) Close() error {
// 	if err := r.channel.Close(); err != nil {
// 		return err
// 	}
// 	return r.conn.Close()
// }
