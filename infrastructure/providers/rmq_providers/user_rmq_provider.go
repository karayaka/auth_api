package rmqproviders

import "github.com/streadway/amqp"

type UserRmqProvider struct {
	channel      *amqp.Channel
	exchangeName string
}

func NewUserProvider(channel *amqp.Channel) *UserRmqProvider {
	ex := "UserEvents"
	channel.ExchangeDeclare(ex, "topic", true, false, false, false, nil)
	return &UserRmqProvider{
		channel:      channel,
		exchangeName: ex,
	}
}

func (up UserRmqProvider) AddMeesageToEvent(name string) error {
	message := amqp.Publishing{
		Body: []byte(name),
	}
	return up.channel.Publish(up.exchangeName, "random-key", false, false, message)
}
