package consumers

import "github.com/streadway/amqp"

func RegisterConsumer(channel *amqp.Channel) {

	go UserEventConsumer(channel)

}
