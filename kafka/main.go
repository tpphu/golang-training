package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	kafkaclient "github.com/uber-go/kafka-client"
	"github.com/uber-go/kafka-client/kafka"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
)

var i = 0

func process(msg kafka.Message) error {
	i++
	if i%3 == 0 {
		return errors.New("oops!")
	}
	return nil
}

func main() {
	// mapping from cluster name to list of broker ip addresses
	brokers := map[string][]string{
		"sample_cluster":     []string{"127.0.0.1:9092"},
		"sample_dlq_cluster": []string{"127.0.0.1:9092"},
	}
	// mapping from topic name to cluster that has that topic
	topicClusterAssignment := map[string][]string{
		"sample_topic": []string{"sample_cluster"},
	}

	// First create the kafkaclient, its the entry point for creating consumers or producers
	// It takes as input a name resolver that knows how to map topic names to broker ip addrs
	client := kafkaclient.New(kafka.NewStaticNameResolver(topicClusterAssignment, brokers), zap.NewNop(), tally.NoopScope)

	// Next, setup the consumer config for consuming from a set of topics
	config := &kafka.ConsumerConfig{
		TopicList: kafka.ConsumerTopicList{
			kafka.ConsumerTopic{ // Consumer Topic is a combination of topic + dead-letter-queue
				Topic: kafka.Topic{ // Each topic is a tuple of (name, clusterName)
					Name:    "sample_topic",
					Cluster: "sample_cluster",
				},
				DLQ: kafka.Topic{
					Name:    "sample_consumer_dlq",
					Cluster: "sample_dlq_cluster",
				},
			},
		},
		GroupName:   "sample_consumer",
		Concurrency: 100, // number of go routines processing messages in parallel
	}
	config.Offsets.Initial.Offset = kafka.OffsetOldest

	// Create the consumer through the previously created client
	consumer, err := client.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	// Finally, start consuming
	if err := consumer.Start(); err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if !ok {
				return // channel closed
			}
			fmt.Println(string(msg.Value()))
			if err := process(msg); err != nil {
				msg.Nack()
			} else {
				msg.Ack()
			}
		case <-sigCh:
			consumer.Stop()
			<-consumer.Closed()
		}
	}
}
