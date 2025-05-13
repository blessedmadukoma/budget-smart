package queue

import (
	"time"

	"github.com/blessedmadukoma/budgetsmart/engine/internal/common"
	"github.com/blessedmadukoma/budgetsmart/engine/pkg/rdb"
)

type Queuer interface {
	Write(common.TaskName, common.QueueName, *Job) error
	Options() QueueOptions
}

type Job struct {
	ID      string        `json:"id"`
	Payload []byte        `json:"payload"`
	Delay   time.Duration `json:"delay"`
	Retry   int           `json:"retry"`
}

type QueueOptions struct {
	Names        map[string]int
	Type         string
	RedisClient  *rdb.Redis
	RedisAddress []string
	// RabbitMQURL    string          // RabbitMQ connection URL
	// RabbitMQClient *rdb.RabbitMQ // RabbitMQ client containing connection and channel

}
