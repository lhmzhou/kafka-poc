package main

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestConnectionToConsumer(t *testing.T) {
	config := &kafka.ConfigMap{"bootstrap.servers": "localhost:8088"}
	p, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	Publish(p, "locate device", "Found")
	x := Next("locate device")
	print(x)
	p.Flush(50 * 1000)
	defer p.Close()

}
