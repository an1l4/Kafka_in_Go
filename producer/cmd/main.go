package main

import (
	kafkaConfig "github.com/an1l4/Kafka_in_Go/config"
	"github.com/an1l4/Kafka_in_Go/producer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC

	producer.Produce(topic, 1000)
}
