package main

import (
	"bufio"
	"context"
	ads "demo_mesh/common-protos/pilot_envoy"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	PORT = "9003"
)

func main() {

    input := bufio.NewScanner(os.Stdin)
    var num  = 0
    for input.Scan() {
		conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}

		defer conn.Close()

		client := ads.NewAggregatedDiscoveryServiceClient(conn)
		stream, err := client.StreamAggregatedResources(context.Background())
		err = stream.Send(&ads.DiscoveryRequest{Discoveryrequest:  strconv.Itoa(num) + ":envoyproxyclient:DiscoveryRequest"})
		num ++
		response,_:= stream.Recv()
		fmt.Printf("envoyproxyclient:received from stream : %v \n", response.GetDiscoveryresponse())
	}

	time.Sleep(time.Second)


}
