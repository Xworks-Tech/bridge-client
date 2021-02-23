package examples

import (
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
	consumer := kChannel.Consume("my-topic22")
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	go func() {
		for elem := range consumer {
			log.Println(string(elem))
		}
	}()

	<-stayAliveFlag

}
