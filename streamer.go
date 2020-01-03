package main

import (
	"time"

	"github.com/Shopify/sarama"
)

func start() {
	var client *kafka.Client // intercept client
	var err error

	for {
		log.Infof("Attempting to connect: %v", kafkaServers)
		client, err = kafka.NewClient(kafkaServers, consumerGroup)
		if err != nil {
			log.Error(err)
			log.Info("Sorry, unable to connect. Waiting...")
			time.Sleep(5 * time.Second)
		}

		if err == nil {
			break
		}
	}

	go stream(client) // stream goroutine

	streamControl <- true // launch

	readControl(client) // start reading from the control topic
}

func getStream(client *kafka.Client, topic string) *kafka.Stream {
	stream, err := client.GetStream(topic)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Connected: %v", topic)
	return stream
}

func readControl(client *kafka.Client) {
	controlStream := getStream(client, controlTopic)

	control, errors := make(chan *sarama.ConsumerMessage), make(chan error)

	go controlStream.Read(control, errors)

	for {
		select {
		case err := <-errors:
			if err != nil {
				log.Fatalf("Error reading from control topic: %v", err)
			}
		case bytes := <-control:
			msg := string(bytes.Value)
			switch msg {
			case "start":
				log.Info("Starting record streamer.")
				streamControl <- true
				break
			case "stop":
				log.Info("Stopping record streamer.")
				streamControl <- false
				break
			default:
				log.Errorf("Sorry, invalid control message: %v", msg)
			}
		}
	}
}

func stream(client *kafka.Client) {
	outputStream := getStream(client, outputTopic)
	inputStream := getStream(client, recordTopic)

	for {
		select {
		case open := <-streamControl:
			if open {
				log.Infof("Reading from record topic: %v", recordTopic)
				inputStream := getStream(client, recordTopic)
				go kafka.Pipe(inputStream, outputStream)
			}

			if !open {
				log.Infof("Closing record topic: %v", recordTopic)
				err := inputStream.Close()
				if err != nil {
					log.Fatalf("Error closing record streamer: %v", err)
				}
			}
		}
	}
}
