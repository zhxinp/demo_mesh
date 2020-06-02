// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package mixs

import (
	"context"
	pb "demo_mesh/common-protos/mixer"
	"demo_mesh/pkg/gtls"
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const (
	PORT = "9091"
	GoPath = "C:/Users/Administrator/go/src/"
)

type MixerServer struct {}

//前置条件检查以及配额分配
func (s *MixerServer) Check (ctx context.Context, cr *pb.CheckRequest)  (*pb.CheckResponse, error) {
	///////////////////////////////////////////////
	//业务逻辑
	///////////////////////////////////////////////
	var checkResponse = pb.CheckResponse{}
	checkResponse01 = strings.ToLower(checkResponse01)
	err := yaml.Unmarshal([]byte(checkResponse01),&checkResponse)
	if err != nil {
		fmt.Errorf("%v" , err)
	}
	checkResponsejson, _ := json.MarshalIndent(checkResponse,"","   ")
	fmt.Printf("\n CheckResponse: \n %s \n",checkResponsejson)
	return &checkResponse, nil
}
//报告遥测数据
func (s *MixerServer) Report (ctx context.Context, rr *pb.ReportRequest) (*pb.ReportResponse, error) {

	//response := &pb.ReportResponse{

	//}
	return &pb.ReportResponse{}, nil
}

func StartServer() {
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
	pb.RegisterMixerServer(server,&MixerServer{}	)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	fmt.Printf("开始监听：%s", PORT)
	server.Serve(lis)

}