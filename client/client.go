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

func (kc *KafkaChannel) Consume(topic string) (<-chan []byte, error) {
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
	return readChan, nil
}

func (kc *KafkaChannel) Produce(topic string) chan<- []byte {
	writeChan := make(chan []byte)

	go func(writer *chan []byte) {
		stream, err := kc.Client.Produce(context.Background())
		if err != nil {
			log.Fatalf("Error creating producer stream: %v", err)
		}
		for item := range *writer {
			stream.Send(&bridge.PublishRequest{
				Topic: topic,
				OptionalContent: &bridge.PublishRequest_Content{
					Content: item,
				},
			})
			response, err := stream.Recv()
			if err != nil {
				stream.CloseSend()
				break

			}

			if response.Success == false {
				switch data := response.Message.(type) {

				case *bridge.ProduceResponse_Content:
					log.Println(data)
					break

				default:
					log.Println("Could not decode response")
					break

				}

			}

		}
		close(*writer)
	}(&writeChan)

	return writeChan
}
