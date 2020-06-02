// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package xds

import (
	"encoding/json"
	"strings"

	//"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestRoute(t *testing.T) {
var routeFile01 = `
Name: "9080"
VirtualHosts:
- Name: "reviews.default.svc.cluster.local:9080"
  Domains:
  - "reviews.default.svc.cluster.local"
  - "reviews.default.svc.cluster.local:9080"
  - "reviews"
  - "reviews:9080"
  - "reviews.default.svc.cluster"
  - "reviews.default.svc.cluster:9080"
  - "reviews.default.svc"
  - "reviews.default.svc:9080"
  - "reviews.default"
  - "reviews.default:9080"
  - "10.97.122.5"
  - "10.97.122.5:9080"
  Routes:
  - Match:
      Prefix: "/"
      Headers:
        Name: "end-user"
        ExactMatch: "jason"
    Route:
      Cluster: "outbound|9080|v2|reviews.default.svc.cluster.local"
    typedPerFilterConfig: 
      EnvoyFault:
        Delay: "2s" 
`

var routeFile02 = `
name: "9080"
virtualhosts:
- name: "reviews.default.svc.cluster.local:9080"
  domains:
  - "reviews.default.svc.cluster.local"
  - "reviews.default.svc.cluster.local:9080"
  - "reviews"
  - "reviews:9080"
  - "reviews.default.svc.cluster"
  - "reviews.default.svc.cluster:9080"
  - "reviews.default.svc"
  - "reviews.default.svc:9080"
  - "reviews.default"
  - "reviews.default:9080"
  - "10.97.122.5"
  - "10.97.122.5:9080"
  routes:
  - match:
      prefix: "/"
      headers:
        name: "end-user"
        exactmatch: "jason"
    route:
      cluster: "outbound|9080|v2|reviews.default.svc.cluster.local"
    typedperfilterconfig: 
      envoyfault:
        delay: "2s" 
`

var routeFile03 = `
Name: "9080"
VirtualHosts:
- Name: "reviews.default.svc.cluster.local:9080"
  Domains:
  - "reviews.default.svc.cluster.local"
  - "reviews.default.svc.cluster.local:9080"
  - "reviews"
  - "reviews:9080"
  - "reviews.default.svc.cluster"
  - "reviews.default.svc.cluster:9080"
  - "reviews.default.svc"
  - "reviews.default.svc:9080"
  - "reviews.default"
  - "reviews.default:9080"
  - "10.97.122.5"
  - "10.97.122.5:9080"
  Routes:
  - Match:
      Prefix: "/"
      Headers:
        Name: "end-user"
        ExactMatch: "jason"
    Route:
      WeightedClusters:
        Clusters:
        - Name: "outbound|9080|v2|reviews.default.svc.cluster.local"
          Weight: 50
        - Name: "outbound|9080|v3|reviews.default.svc.cluster.local"
          Weight: 50
      Cluster: "outbound|9080|v2|reviews.default.svc.cluster.local"
    typedPerFilterConfig: 
      EnvoyFault:
        Delay: "2s" 
`



routeFile01 = strings.ToLower(routeFile03)

fmt.Printf("%s \n",routeFile02)
var route02 = Route{}
err := yaml.Unmarshal([]byte(routeFile01),&route02)
if err != nil {
	t.Errorf("%v" , err)
}
route02json, _ := json.MarshalIndent(route02,"","   ")
fmt.Printf("%s \n",route02json)
}