package examples

import (
	"log"
	"time"

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
	kChannel := client.KafkaChannel{Stream: bridge.NewKafkaStreamClient(cc)}
	consumer, producer, err := kChannel.SubscribeToTopic("my-topic")
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}
	stayAliveFlag := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			producer <- []byte("hello!")
		}
	}()

	go func() {
		for elem := range consumer {
			log.Println(string(elem))
		}
	}()

	<-stayAliveFlag

}
