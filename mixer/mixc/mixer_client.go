package mixc

import (
	"context"
	pb "demo_mesh/common-protos/mixer"
	"demo_mesh/pkg/gtls"
	"google.golang.org/grpc"
	"log"
)

const (
	PORT = "9091"
	GoPath = "C:/Users/Administrator/go/src/"
)

func MixerClient(request pb.CheckRequest) (*pb.CheckResponse, error) {
	tlsClient := gtls.Client{
		ServerName: "go-grpc-example",
		CaFile:     GoPath + "demo_mesh/common/certificate-authority/ca.pem",
		CertFile:   GoPath + "demo_mesh/common/certificate-authority/client/client.pem",
		KeyFile:    GoPath + "demo_mesh/common/certificate-authority/client/client.key",
	}
	c, err := tlsClient.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()
    client := pb.NewMixerClient(conn)
	checkResponse, err := client.Check(context.Background(),&request)
	//fmt.Printf("%s \n",checkResponse )
	if err != nil {
		return nil, err
	}
	return checkResponse,err

}
