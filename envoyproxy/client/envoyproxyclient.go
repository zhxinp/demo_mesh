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
	"demo_mesh/pkg/gtls"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

const (
	PORT = "9002"
	GoPath = "C:/Users/Administrator/go/src/"
)

func main() {
	tlsClient := gtls.Client{
		ServerName: "go-grpc-example",
		CaFile:   GoPath+"demo_mesh/common/certificate-authority/ca.pem",
		CertFile: GoPath+"demo_mesh/common/certificate-authority/client/client.pem",
		KeyFile:  GoPath+"demo_mesh/common/certificate-authority/client/client.key",
	}
	c, err := tlsClient.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}
	var num  = 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}

		defer conn.Close()

		client := ads.NewAggregatedDiscoveryServiceClient(conn)
		stream, err := client.StreamAggregatedResources(context.Background())
		err = stream.Send(&ads.DiscoveryRequest{Discoveryrequest:strconv.Itoa(num)+"client:DiscoveryRequest"})
		num ++
		response,err:= stream.Recv()
		if err != nil {
			fmt.Errorf("err : %v", err)
		}
		fmt.Printf("client:received from stream : %v \n", response.GetDiscoveryresponse())
	}



}