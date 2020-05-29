

package xds
import "encoding/json"

type HttpConnectionManager struct {

}

type Listener struct {
	Name         string //virtualInbound
	Address      *AddressEntries
	FilterChains *FilterChainsEntries
}
func (lis *Listener)GetName() string {
	if lis !=nil {
		return lis.Name
	}
	return ""
}

func (lis *Listener)GetAddress() *AddressEntries {
	if lis !=nil {
		return lis.Address
	}
	return nil
}

func (lis *Listener)GetFilterChains() *FilterChainsEntries {
	if lis !=nil {
		return lis.FilterChains
	}
	return nil
}
type AddressEntries struct {
	SocketAddress string  //0.0.0.0
	PortValue int32  //15006
}

func (ad *AddressEntries) GetSocketAddress() string {
	if ad != nil {
		return ad.SocketAddress
	}
	return ""
}

func  (ad *AddressEntries) GetPortValue() int32 {
	if ad != nil {
		return ad.PortValue
	}
	return 0
}

type FilterChainsEntries struct {
	FilterChainMatch *FilterChainMatchEntries //
	Filters *FiltersEntries
}

func (fi *FilterChainsEntries) GetFilterChainMatch() *FilterChainMatchEntries {
	if fi != nil {
		return fi.FilterChainMatch
	}
	return nil
}

func (fi *FilterChainsEntries) GetFilters() *FiltersEntries {
	if fi != nil {
		return fi.Filters
	}
	return nil
}

type FilterChainMatchEntries struct {
	DestinationPort int32 //9080
}

func (fi *FilterChainMatchEntries)GetDestinationPort() int32 {
	if fi != nil {
		return fi.DestinationPort
	}
	return 0
}

type FiltersEntries struct {
	Name string //envoy.http_connection_manager
	Rds *RdsEntries
	HttpFilters *HttpFiltersEntries
}

func (fi *FiltersEntries) GetName() string {
	if fi != nil {
		return fi.Name
	}
	return ""
}

func (fi *FiltersEntries) GetHttpFilters() *HttpFiltersEntries {
	if fi != nil {
		return fi.HttpFilters
	}
	return nil
}

type RdsEntries struct {
	RouteConfigName string   //9080
}

func (rd *RdsEntries) GetRouteConfigName() string {
	if rd != nil {
		return rd.RouteConfigName
	}
	return ""
}
type HttpFiltersEntries struct {
	Name string  //mixer
	CheckCluster string  //outbound|9091||istio-policy.istio-system.svc.cluster.local
	ReportCluster string  //outbound|9091||istio-telemetry.istio-system.svc.cluster.local
}

func (ht *HttpFiltersEntries) GetName() string {
	if ht != nil {
		return ht.Name
	}
	return ""
}

func (ht *HttpFiltersEntries) GetCheckCluster() string {
	if ht != nil {
		return ht.CheckCluster
	}
	return ""
}

func (ht *HttpFiltersEntries) GetReportCluster() string {
	if ht != nil {
		return ht.ReportCluster
	}
	return ""
}

func (lis *Listener) GOtoJson() ([]byte, error ){
	listenerJson , err:= json.MarshalIndent(*lis, "", "   ")
	return listenerJson, err
}



func DefaultListener() *Listener {
	return &Listener{
		Name:  "virtualInbound",
		Address: DefaultAddressEntries(),
		FilterChains: DefaultFilterChainsEntries(),
	}
}

func DefaultAddressEntries() *AddressEntries {
	return &AddressEntries{
		SocketAddress: "0.0.0.0",
		PortValue: 15006,
	}
}

func DefaultFilterChainsEntries() *FilterChainsEntries {
	return &FilterChainsEntries{
		FilterChainMatch: DefaultFilterChainMatchEntries(),
		Filters: DefaultFiltersEntries(),


	}
}

func DefaultFilterChainMatchEntries() *FilterChainMatchEntries {
	return &FilterChainMatchEntries{
		DestinationPort: 9080,
	}
}

func DefaultFiltersEntries() *FiltersEntries {
	return &FiltersEntries{
		Name:"envoy.http_connection_manager",
		Rds:DefaultRdsEntries(),
		HttpFilters: DefaultHttpFiltersEntries(),
	}
}

func DefaultHttpFiltersEntries() *HttpFiltersEntries{
	return &HttpFiltersEntries{
		Name:          "mixer",
		CheckCluster:  "outbound|9091||istio-policy.istio-system.svc.cluster.local",
		ReportCluster: "outbound|9091||istio-telemetry.istio-system.svc.cluster.local",
	}
}

func DefaultRdsEntries () *RdsEntries {
	return &RdsEntries{
		RouteConfigName: "9080",
	}
}
