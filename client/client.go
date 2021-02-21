package client

import (
	"context"
	"log"

	bridge "github.com/Xworks-Tech/bridge-client/proto"
)

// KafkaChannel
// A bi-directional streaming channel to listen and publish to a kafka topic
type KafkaChannel struct {
	Stream bridge.KafkaStreamClient
}

// SubscribeToTopic
// Subscribes to the kafka topic through the rpc call and returns a write and read channel
func (kc *KafkaChannel) SubscribeToTopic(topic string) (<-chan []byte, chan<- []byte, error) {
	stream, err := kc.Stream.Subscribe(context.Background())
	if err != nil {
		return nil, nil, err
	}

	write, read := make(chan []byte), make(chan []byte)

	//
	go func(st *bridge.KafkaStream_SubscribeClient, reader *chan []byte) {
		for item := range *reader {
			if err := (*st).Send(&bridge.PublishRequest{
				Topic: topic,
				OptionalContent: &bridge.PublishRequest_Content{
					Content: item,
				},
			}); err != nil {
				log.Printf("Error sending message to bridge: %v", err)
				(*st).CloseSend()
				close(*reader)
				break
			}
		}
	}(&stream, &read)

	go func(st *bridge.KafkaStream_SubscribeClient, writer *chan []byte) {
		for {
			response, err := (*st).Recv()
			if err != nil {
				log.Printf("Error receiving a response from bridge: %v", err)
				close(*writer)
				break
			}
			switch data := response.OptionalContent.(type) {

			case *bridge.KafkaResponse_Content:
				*writer <- *&data.Content
				break

			default:
				break

			}
		}
	}(&stream, &write)

	return write, read, nil
}
