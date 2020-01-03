package main

import (
	"fmt"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestConnectToProducer(t *testing.T) {
	actual, err := ConnectToProducer()
	if err != nil {
		panic(err)
	}
	cfg := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
	expected, err := kafka.NewProducer(cfg)
	if err != nil {
		panic(err)
	}

	if actual == expected {
		fmt.Print(actual, expected)
		t.Error("The connection string do not match")
	}
}
