package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

// Start triggers a new producer routine, creates the underlying Kafka producer client and events receiver, leaving it ready to handle events.
func Start(topics string, msg string) {

	sigChannel := make(chan bool)
	defer close(sigChannel)

	// producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kakfa:8088"}) // create a new producer and connects to Kafka cluster
	if err != nil {
		// log.Fatal(err)
		log.WithError(err).Error("Failed to create kafka producer. Please try again.")
		// panic(err)
		return err
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
					// fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
					sigChannel <- false
				} else {
					// log.Printf("Produced message to %v\n", m.TopicPartition)
					log.Printf("Produced message to %v\n", string(m.Value))
					sigChannel <- true
				}
				return
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topics, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)

	// return <- sig

	// Wait for message deliveries
	producer.Flush(15 * 1000)
	println("Halting producer...")
}
