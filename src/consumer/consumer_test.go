package main

import (
	"fmt"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestConnectionToProducer(t *testing.T) {
	actual, err := ConnectToProducer()
	if err != nil {
		panic(err)
	}
	config := &kafka.ConfigMap{"bootstrap.servers": "localhost:8088"}
	expected, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	if actual == expected {
		fmt.Print(actual, expected)
		t.Error("The connection string do not match")
	}
}
