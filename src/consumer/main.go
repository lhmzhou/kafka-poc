package consumer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Start triggers a new consumer routine
func Start(topic string) {

	// logger := util.GetLogger("consumer", "KafkaConsumer::dispatcher")
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":           "pocGroup",
		"enable.auto.commit": "false",
		"auto.offset.reset":  "earliest",
	})
	failOnError(err, "Consumer cannot be created")

	file := createFile()
	defer file.Close()
	// consumer.SubscribeTopics([]string{"pocTopic", "^aRegex.*[Tt]opic"}, nil)

	c.SubscribeTopics([]string{topic}, nil)

	log.Printf("Trying to receive incoming message...")

	for {
		// -1 for indefinite wait or timeout
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			file.WriteString(string(msg.Value) + "\n")
			fmt.Printf("Saved on file %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
}
