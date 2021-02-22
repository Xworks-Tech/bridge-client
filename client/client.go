package client

import (
	"context"
	"log"

	bridge "github.com/Xworks-Tech/bridge-client/proto"
	"google.golang.org/grpc"
)

// KafkaChannel
// A bi-directional streaming channel to listen and publish to a kafka topic
func New(cc *grpc.ClientConn) KafkaChannel {
	client := bridge.NewKafkaStreamClient(cc)
	stream, err := client.Subscribe(context.Background())
	if err != nil {
		log.Fatalf("Error setting up stream: %v", err)
	}
	return KafkaChannel{
		Client: client,
		Stream: stream,
	}
}

type KafkaChannel struct {
	Client bridge.KafkaStreamClient
	Stream bridge.KafkaStream_SubscribeClient
}

// SubscribeToTopic
// Subscribes to the kafka topic through the rpc call and returns a write and read channel
func (kc *KafkaChannel) SubscribeToTopic(topic string) (<-chan []byte, chan<- []byte, error) {

	write, read := make(chan []byte), make(chan []byte)

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
	}(&kc.Stream, &read)

	go func(st *bridge.KafkaStream_SubscribeClient, writer *chan []byte) {
		for {
			response, err := (*st).Recv()
			log.Println("Received message")
			if err != nil {
				log.Printf("Error receiving a response from bridge: %v", err)
				close(*writer)
				break
			}
			switch data := response.OptionalContent.(type) {
			case *bridge.KafkaResponse_Content:
				*writer <- *&data.Content
			default:
				break

			}
		}
	}(&kc.Stream, &write)

	return write, read, nil
}
