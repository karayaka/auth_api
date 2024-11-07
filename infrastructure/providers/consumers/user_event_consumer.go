package consumers

import (
	"fmt"

	"github.com/streadway/amqp"
)

func UserEventConsumer(channel *amqp.Channel) error {

	///Cunsumere ait quyruk oluşturuluyor
	channel.QueueDeclare("UserEventQueue", true, false, false, false, nil)

	//kuyruk ile event bind ediliyor
	channel.QueueBind("UserEventQueue", "#", "UserEvents", false, nil)
	//kuyruk dinleniyor
	msgs, err := channel.Consume("UserEventQueue", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever

	return nil
}
