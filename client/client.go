package client

import (
	"context"
	"log"

	bridge "github.com/Xworks-Tech/bridge-client/proto"
)

// KafkaChannel

type KafkaChannel struct {
	Client bridge.KafkaStreamClient
}

func (kc *KafkaChannel) Consume(topic string) <-chan []byte {
	readChan := make(chan []byte)
	go func(reader *chan []byte) {
		stream, err := kc.Client.Consume(context.Background(), &bridge.ConsumeRequest{
			Topic: topic,
		})
		if err != nil {
			log.Fatalf("Error creating cosumer stream: %v", err)
		}
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error creating cosumer stream: %v", err)
			}
			switch data := response.OptionalContent.(type) {
			case *bridge.KafkaResponse_Content:
				*reader <- *&data.Content
			default:
				break

			}
		}
	}(&readChan)
	return readChan
}

func (kc *KafkaChannel) Produce(topic string) chan<- []byte {
	writeChan := make(chan []byte)

	go func(writer *chan []byte) {
		defer close(*writer)

		stream, err := kc.Client.Produce(context.Background())
		if err != nil {
			log.Fatalf("Error creating producer stream: %v", err)
		}
		for item := range *writer {
			err := stream.Send(&bridge.PublishRequest{
				Topic: topic,
				OptionalContent: &bridge.PublishRequest_Content{
					Content: item,
				},
			})
			if err != nil {
				log.Printf("Error sending message to bridge: %v", err)
				stream.CloseSend()
				break

			}
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving from ack from bridge: %v", err)
				stream.CloseSend()
				break

			}

			if response.Success == false {
				log.Printf("Unsuccessful request sent to bridge: %s", string(item))
				switch data := response.Message.(type) {

				case *bridge.ProduceResponse_Content:
					log.Printf("Error writing to bridge: %v", data)

				default:
					log.Println("Could not decode response")

				}

			}

		}

	}(&writeChan)

	return writeChan
}

func (kc *KafkaChannel) MultiProduce() <-chan []string {
	requestChan := make(chan []string)
	go func(writer *chan []string) {
		defer close(*writer)

		stream, err := kc.Client.Produce(context.Background())
		if err != nil {
			log.Fatalf("Error creating producer stream: %v", err)
		}
		for item := range *writer {
			err := stream.Send(&bridge.PublishRequest{
				Topic: item[0],
				OptionalContent: &bridge.PublishRequest_Content{
					Content: []byte(item[1]),
				},
			})
			if err != nil {
				log.Printf("Error sending message to bridge: %v", err)
				stream.CloseSend()
				break

			}
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving from ack from bridge: %v", err)
				stream.CloseSend()
				break

			}

			if response.Success == false {
				log.Printf("Unsuccessful request sent to bridge: ")
				switch data := response.Message.(type) {

				case *bridge.ProduceResponse_Content:
					log.Printf("Error writing to bridge: %v", data)

				default:
					log.Println("Could not decode response")

				}

			}

		}

	}(&requestChan)
	return requestChan
}
