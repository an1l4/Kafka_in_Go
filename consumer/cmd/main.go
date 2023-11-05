package main

import (
	kafkaConfig "github.com/an1l4/Kafka_in_Go/config"
	"github.com/an1l4/Kafka_in_Go/consumer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC

	consumer.Consume(topic)
}
