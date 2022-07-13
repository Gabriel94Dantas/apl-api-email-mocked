package contexts

import (
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConsumer() *confluentKafka.Consumer {
	var p *confluentKafka.Consumer
	var err error = nil

	if os.Getenv("BROKER_HOST") == "" {
		os.Setenv("BROKER_HOST", "localhost")
	}

	p, err = confluentKafka.NewConsumer(&confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BROKER_HOST"),
		"group.id":          "Mygroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return p
}
