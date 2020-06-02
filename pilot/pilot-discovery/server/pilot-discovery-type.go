// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package main

type Environment struct {
	ServiceDiscovery
	IstioConfigStore
}

//列出service和instances
type ServiceDiscovery interface {
	Services() ([] *Service,error)
	GetService(hostname string)( *Service,error)
}
//列出路由规则
type IstioConfigStore interface {
	ConfigStore
}

type Service struct {
   Name string
   Namespace string
   Hostname string
   address string
}

//参照istio/pilot/pkg/proxy/envoy/v2/config.pb.go
type MeshConfig struct {
	//向mixer发送Check请求的地址
	MixerCheckServer string
	//向Mixer返送Report的地址
	MixerReportServer string
	ProxyListenPort int32

}

//参照istio/pilot/pkg/proxy/envoy/v2
type ProxyConfig struct {
	//生成的配置文件存放位置
	ConfigPath string
	//Envoy可执行文件的路径
	BinaryPath string
    //服务发现地址,可根据此地址获取xDS
	DiscoveryAddress string

}

//ConfigMega中的数据用于产生FQDN
//pilot/pkg/model/config.go
type Config struct {
	Name string
	Namespace string
	Type string
	Version string
	Domain string


}

type XdsConnection struct {
	//envoy客户端地址
	PeerAddr string
	pushChannel chan *XdsEvent



}

type XdsEvent struct {
	push *PushContext
}

type PushContext struct {
	Services []*Service
	Mesh *MeshConfig
	VirtualServices []Config
	QuotaSpec []Config
	Version string
}