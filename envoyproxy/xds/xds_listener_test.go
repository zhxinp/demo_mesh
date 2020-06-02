// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package xds


import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGotoJson(t *testing.T) {
	listener := DefaultListener()
	fmt.Printf("%T",listener)
	datajson, _ := json.MarshalIndent(listener, "", "   ")
	fmt.Printf("%s \n", datajson)
	datayaml, _ := yaml.Marshal(listener)
	fmt.Printf("%s \n", datayaml)



}

func TestYamlToGo(t *testing.T) {
yamlFile := `
name: virtualInbound
address:
  socketaddress: 0.0.0.0
  portvalue: 15006
filterchains:
  filterchainmatch:
    destinationport: 9080
  filters:
    name: envoy.http_connection_manager
    rds:
      routeconfigname: "9080"
    httpfilters:
      name: mixer
      checkcluster: outbound|9091||istio-policy.istio-system.svc.cluster.local
      reportcluster: outbound|9091||istio-telemetry.istio-system.svc.cluster.local
`
goFile := Listener{}
yaml.Unmarshal([]byte(yamlFile),&goFile)
fmt.Print(goFile.GetAddress().GetSocketAddress())
//datajson, _ := json.MarshalIndent(goFile, "", "   ")
//fmt.Printf("%s \n", datajson)
}
