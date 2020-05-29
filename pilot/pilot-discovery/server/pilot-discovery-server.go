package main

import (
	ads "demo_mesh/common-protos/pilot_envoy"
	"demo_mesh/pkg/gtls"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type DiscoveryServer struct {
	//Env *Environment

}

//protoc -I .  proto/stream.proto --go_out=plugins=grpc:proto

const (
	PORT = "9002"
	GoPath = "C:/Users/Administrator/go/src/"
)

func main() {
	tlsServer := gtls.Server{
		CaFile:   GoPath+"demo_mesh/common/certificate-authority/ca.pem",
		CertFile: GoPath+"demo_mesh/common/certificate-authority/server/server.pem",
		KeyFile:  GoPath+"demo_mesh/common/certificate-authority/server/server.key",
	}
	c, err := tlsServer.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
	}
	server := grpc.NewServer(opts...)
	ads.RegisterAggregatedDiscoveryServiceServer(server, &DiscoveryServer{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	fmt.Printf("开始监听：%s", PORT)
	server.Serve(lis)

}

func (s *DiscoveryServer) StreamAggregatedResources(stream ads.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	//reqChannel := make(chan *ads.DiscoveryRequest)
	request,err := stream.Recv()
	//客户端发来的请求存储在reqChannel里面
	//reqChannel <- request
	//获取最新的信息

	if err != nil {
		fmt.Errorf("err : %v", err)
	}
	fmt.Printf("server:received request from strem: %v\n", request.GetDiscoveryrequest())
	stream.Send(&ads.DiscoveryResponse{Discoveryresponse:"server:discoveryResponse"})
	return nil
}