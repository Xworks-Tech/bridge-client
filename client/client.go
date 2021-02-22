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

func (kc *KafkaChannel) Subscribe(topic string) (<-chan []byte, error) {
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
