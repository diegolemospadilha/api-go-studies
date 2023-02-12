package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChannel chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "api-go",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChannel <- msg
		}
	}
}
