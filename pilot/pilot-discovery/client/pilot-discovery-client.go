// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
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
