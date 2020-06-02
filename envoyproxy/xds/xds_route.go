// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package xds

type Route struct {
	Name string //9080
	VirtualHosts []*VirtualHostsEntries
}


type VirtualHostsEntries struct {
	Name string //
	Domains [] string
	Routes *[] RoutesEntries
}

type RoutesEntries struct {
	Match *MatchEntries
	Route *RouteEntries
	TypedPerFilterConfig *TypedPerFilterConfigEntries
}

type MatchEntries struct {
	Prefix string //"/"
	Headers *HeadersEntries
}

type HeadersEntries struct {
	Name string  //"end-user"
	ExactMatch string //"jason"

}

type RouteEntries struct {
    Cluster string
	WeightedClusters *WeightedClustersEntries

}

type WeightedClustersEntries struct {
	Clusters *[] ClusterEntries
}

type ClusterEntries struct {
	Name string
	Weight int8
	TypedPerFilterConfig *TypedPerFilterConfigEntries
}



type TypedPerFilterConfigEntries struct {
	EnvoyFault *EnvoyFaultEntries
}

type EnvoyFaultEntries struct {
	Delay string //"2s"
	Abort *AbortEntries
}

type AbortEntries struct {
	HttpStatus int8 //500
}

