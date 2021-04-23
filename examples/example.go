package examples

import (
	"context"
	"log"

	"github.com/Xworks-Tech/bridge-client/client"
	bridge "github.com/Xworks-Tech/bridge-client/proto"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error starting grpc client: %v", err)
	}

	defer cc.Close()
	kChannel := client.KafkaChannel{
		Client: bridge.NewKafkaStreamClient(cc),
	}
	stayAliveFlag := make(chan bool)
	consumer, consumerCallback := kChannel.Consume("my-topic", "my-id")
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	go func() {
		for elem := range consumer {
			log.Println(string(elem))
		}
	}()

	// Produce asynchronously
	writer, writeClose := kChannel.AsyncProduce("my-topic")
	writer <- []byte("My message")

	<-stayAliveFlag // Stalls the program indefinetly

	if err := consumerCallback(); err != nil {
		log.Panicf("Error closing consumer stream")
	} // Used to close consumer stream

	if err := writeClose(); err != nil {
		log.Panicf("Error closing writer: %v", err)
	} // Used to close write stream

}

func synchronousProduce() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error starting grpc client: %v", err)
	}

	defer cc.Close()
	kChannel := client.KafkaChannel{
		Client: bridge.NewKafkaStreamClient(cc),
	}
	writer, err := kChannel.CreateWriter("my-topic")
	if err != nil {
		log.Printf("Error creating writer: %v", err)
	}
	msg := "my message"
	if err := writer.Produce(msg); err != nil {
		log.Printf("Error writing message (%s): %v", msg, err)
	}
}

func rawClientConsume(topic string, ID string) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error starting grpc client: %v", err)
	}

	defer cc.Close()
	kChannel := client.KafkaChannel{
		Client: bridge.NewKafkaStreamClient(cc),
	}
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := kChannel.Client.Consume(ctx, &bridge.ConsumeRequest{
		Topic: topic,
		Id:    ID,
	})
	closeCallback := func() error {
		err := stream.CloseSend()
		cancel()
		return err
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			break
		}
		switch data := response.OptionalContent.(type) {
		case *bridge.KafkaResponse_Content:
			log.Printf("Received message from bridge: %s", string(*&data.Content))
		default:
			break

		}
	}
	closeCallback()
}
