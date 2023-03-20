package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Set up a signal handler to gracefully shut down the application
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Read the Kafka brokers from the environment variable
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		log.Fatal("KAFKA_BROKERS environment variable not set")
	}

	// Read the JSON file
	jsonFile, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("error reading JSON file: %v", err)
	}

	// Create a Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokers},
		Topic:   "test-topic",
	})

	// Start a goroutine to consume messages from Kafka
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			// Read a message from Kafka
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("error reading message: %v", err)
				continue
			}

			// Determine the message format (protobuf or JSON)
			var message interface{}
			contentType := ""
			for _, header := range msg.Headers {
				if header.Key == "content-type" {
					contentType = string(header.Value)
					break
				}
			}
			switch contentType {
			case "application/protobuf":
				message = &MyProtoMessage{}
				err = proto.Unmarshal(msg.Value, message.(proto.Message))
			case "":
				message = &MyJSONMessage{}
				err = json.Unmarshal(jsonData, message)
			default:
				log.Printf("unknown content-type header: %s", contentType)
				continue
			}

			if err != nil {
				log.Printf("error unmarshaling message: %v", err)
				continue
			}

			// Process the message
			switch m := message.(type) {
			case *MyProtoMessage:
				processProtoMessage(m)
			case *MyJSONMessage:
				processJSONMessage(m)
			}
		}
	}()

	// Wait for a signal to shut down the application
	sig := <-sigChan
	log.Printf("received signal %s, shutting down", sig)

	// Close the Kafka reader
	err = reader.Close()
	if err != nil {
		log.Printf("error closing Kafka reader: %v", err)
	}

	// Wait for the goroutines to finish
	wg.Wait()
}

type MyProtoMessage struct {
    Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
    Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}


type MyJSONMessage struct {
    Id      int32  `json:"id"`
    Message string `json:"message"`
}


func processProtoMessage(m *MyProtoMessage) {
	// process protobuf message here
	fmt.Printf("received protobuf message: %+v\n", m)
}

func processJSONMessage(m *MyJSONMessage) {
	// process JSON message here
	fmt.Printf("received JSON message: %+v\n", m)
}
