package rmqproviders

import "github.com/streadway/amqp"

type RmqProvider struct {
	UserRmqProvider UserRmqProvider
}

func NewRmqProvider(channel *amqp.Channel) *RmqProvider {
	return &RmqProvider{
		UserRmqProvider: *NewUserProvider(channel),
	}
}
