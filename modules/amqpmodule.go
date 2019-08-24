package modules

import "github.com/streadway/amqp"

var AMQPModule *amqp.Connection

func InitAMQPModule() {
	AMQPModule = amqp.Dial
}
