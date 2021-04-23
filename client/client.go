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

func (kc *KafkaChannel) Consume(topic string, consumerID string) (<-chan []byte, func() error) {
	readChan := make(chan []byte)
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := kc.Client.Consume(ctx, &bridge.ConsumeRequest{
		Topic: topic,
		Id:    consumerID,
	})
	if err != nil {
		log.Fatalf("Error creating cosumer stream: %v", err)
	}
	go func(reader *chan []byte, consumeStream *bridge.KafkaStream_ConsumeClient) {
		defer func() {
			if recover() != nil {
				log.Println("Consumer routine closed")
			}
		}()
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error creating cosumer stream: %v", err)
				break
			}
			switch data := response.OptionalContent.(type) {
			case *bridge.KafkaResponse_Content:
				*reader <- *&data.Content
			default:
				break

			}
		}

	}(&readChan, &stream)
	closeCallback := func() error {
		err := stream.CloseSend()
		close(readChan)
		cancel()
		return err
	}
	return readChan, closeCallback
}

func (kc *KafkaChannel) Produce(topic string) chan<- []byte {
	writeChan := make(chan []byte)

	go func(writer *chan []byte) {
		stream, err := kc.Client.Produce(context.Background())
		if err != nil {
			log.Fatalf("Error creating producer stream: %v", err)
		}

		defer close(*writer)
		defer stream.CloseSend()

		for item := range *writer {
			err := stream.Send(&bridge.PublishRequest{
				Topic: topic,
				OptionalContent: &bridge.PublishRequest_Content{
					Content: item,
				},
			})
			if err != nil {
				log.Printf("Error sending message to bridge: %v", err)
				break

			}
			response, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving from ack from bridge: %v", err)
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
				break

			}

		}

	}(&writeChan)

	return writeChan
}
