// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package main

import (
	xds "demo_mesh/envoyproxy/xds"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGotoJson(t *testing.T) {
	listener := xds.DefaultListener()

	datajson, _:= json.MarshalIndent(listener,"","   ")
	fmt.Printf("%s \n", datajson)
	datayaml, _ := yaml.Marshal(listener)
	fmt.Printf("%s \n", datayaml)

	//listener.FilterChains.Filters.HttpFilters.CheckCluster = "hello"
	//fmt.Println(listener.GetFilterChains().GetFilters().GetHttpFilters().CheckCluster)
	/*
	listener := Listener
	{Name:"virtualInbound",Address:
		AddressEntries{
		   SocketAddress: "0.0.0.0",
		   PortValue:     15006,
	}}
	listener.name = "virtualInbound"
	address := AddressEntries{}
	address.socketAddress =  "0.0.0.0"
	address.portValue = 15006
	listener.address = address

	//filtersfiltersEntries{}

	//listener.address.socketAddress = "0.0.0.0"
	//listener.address.portValue = 15006
	listener.filterChains.filterChainMatch.destinationPort = 9080
	listener.filterChains.filters.name = "envoy.http_connection_manager"
	listener.filterChains.filters.httpFilters.name = "mixer"
	listener.filterChains.filters.httpFilters.checkCluster = "outbound|9091||istio-policy.istio-system.svc.cluster.local"
	listener.filterChains.filters.httpFilters.reportCluster = "outbound|9091||istio-telemetry.istio-system.svc.cluster.local"

	fmt.Printf("%v \n", listener)
	data, _ := listener.GOtoJson()

	fmt.Printf("%v \n", data)*/

}

