package client

import (
	"context"
	"errors"
	"fmt"
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

func (kc *KafkaChannel) AsyncProduce(topic string) (chan<- []byte, func() error) {
	writeChan := make(chan []byte)
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := kc.Client.Produce(ctx)
	if err != nil {
		log.Fatalf("Error creating producer stream: %v", err)
	}
	closeCallback := func() error {
		err := stream.CloseSend()
		close(writeChan)
		cancel()
		return err
	}
	go func(writer *chan []byte) {
		defer closeCallback()
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

	return writeChan, closeCallback
}

func (kc *KafkaChannel) CreateWriter(topic string) (*KafkaWriter, error) {
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := kc.Client.Produce(ctx)
	if err != nil {
		cancel()
		return nil, err
	}
	closeCallback := func() error {
		err := stream.CloseSend()
		cancel()
		return err
	}
	return &KafkaWriter{
		Topic:  topic,
		Close:  closeCallback,
		stream: stream,
	}, nil
}

type KafkaWriter struct {
	Topic  string
	Close  func() error
	stream bridge.KafkaStream_ProduceClient
}

func (kw KafkaWriter) Produce(message string) error {
	err := kw.stream.Send(&bridge.PublishRequest{
		Topic: kw.Topic,
		OptionalContent: &bridge.PublishRequest_Content{
			Content: []byte(message),
		},
	})
	if err != nil {
		log.Printf("Error sending message to bridge: %v", err)
		return err

	}
	response, err := kw.stream.Recv()
	if err != nil {
		log.Printf("Error receiving from ack from bridge: %v", err)
		return err

	}

	if response.Success == false {
		log.Printf("Unsuccessful request sent to bridge: %s", message)
		switch data := response.Message.(type) {

		case *bridge.ProduceResponse_Content:
			return errors.New(fmt.Sprintf("Error writing to bridge: %v", data))

		default:
			return errors.New("Could not decode response")

		}

	}
	return nil
}
