package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func mainConsumer(partition int32) {
	kafka := newKafkaConsumer()
	defer kafka.Close()

	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}

func consumeEvents(consumer sarama.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	var CinemaAccount *CinemaAccount
	var err error

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Printf("Kafka error: %s\n", err)
		case msg := <-consumer.Messages():
			msgVal = msg.Value

			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))

				switch logType {
				case "CreateEvent":
					event := new(CreateEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						CinemaAccount, err = event.Process()
					}
				case "PaymentEvent":
					event := new(PaymentEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						CinemaAccount, err = event.Process()
					}
				default:
					fmt.Println("Unknown command: ", logType)
				}

				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Printf("%+v\n\n", *CinemaAccount)
				}
			}
		}
	}
}
