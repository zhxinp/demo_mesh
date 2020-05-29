package main


import (
	ads "demo_mesh/common-protos/pilot_envoy"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type EnvoyProxyServer struct {

}

//protoc -I .  proto/stream.proto --go_out=plugins=grpc:proto




const (
	PORT = "9003"
)

func main() {
	server := grpc.NewServer()
	ads.RegisterAggregatedDiscoveryServiceServer(server, &EnvoyProxyServer{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

func (s *EnvoyProxyServer) StreamAggregatedResources(stream ads.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	request,err := stream.Recv()
	if err != nil {
		fmt.Errorf("err : %v", err)
	}
	fmt.Printf("envoyproxyserver:received request from strem: %v\n", request.GetDiscoveryrequest())
	stream.Send(&ads.DiscoveryResponse{Discoveryresponse:"envoyproxyserver:discoveryResponse"})
	return nil
}
