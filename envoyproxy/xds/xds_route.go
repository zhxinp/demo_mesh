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

