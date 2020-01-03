# kafka-poc

To get up and running, install a Kafka cluster using containers. It's conveniently portable and lightweight, so we can use this on any machine running Docker. Key features to this PoC are 1) create a listener on one or multiples topics 2) create a sarama producer. 

## Usage

Build image from Dockerfile
```
docker build --tag kafka_i /kafka-poc/
```

Install Sarama client
```
$ go get github.com/Shopify/sarama
```

Build producer & consumer
```
$ make build
```

Invocation
```
$ kafka-poc -topic=test -brokers=localhost:8088
```

## References and other supportive links

Sarama is an MIT-licensed Go client library for [Apache Kafka](https://kafka.apache.org/) version 0.8 (and later).

[GoDocs](https://godoc.org/github.com/Shopify/sarama)

[Kafka Protocol Specification](https://cwiki.apache.org/confluence/display/KAFKA/A+Guide+To+The+Kafka+Protocol)

## Contributing 
Pull requests are welcomed with arms wide open.

## On the horizon
[ ] consuming messages from Kafka using sarama-cluster
[ ] intercept and mutate records received by the consumer
[ ] compare stability of Kafka implementation with Rust