// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package mixc

import (
	pb "demo_mesh/common-protos/mixer"
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"strings"
	"testing"
)

var checkRequest00 = `
Attributes:
- Name: "test-attribute-name"
  Value:
    bool_value: true
- Name: "test-attribute-string"
  Value:
    string_value: "test-string"
- Name: "test-attribute-int"
  Value:
    int64_value: 234
- Name: "test-attribute-double"
  Value:
    double_value: 12.3
Quotas:
- Name: "quota-name"
  Amount: 20
  BestEffort: true
`
var checkRequest01 = `
Attributes:
- Name: "test-attribute-name"
  Value:
      BoolValue: true
- Name: "test-attribute-string"
  Value:
      StringValue: "test-string"
- Name: "test-attribute-int"
  Value:
      Int64Value: 234
- Name: "test-attribute-double"
  Value:
      DoubleValue: 12.3
Quotas:
- Name: "quota-name"
  Amount: 20
  BestEffort: true
`
var checkRequest02 = `
Attributes:
- Name: "test-attribute-name"
  bool_value: true
- Name: "test-attribute-string"
  string_value: "test-string"
- Name: "test-attribute-int"
  int64_value: 234
- Name: "test-attribute-double"
  double_value: 12.3
Quotas:
- Name: "quota-name"
  Amount: 20
  BestEffort: true
`

var checkRequest03 = `
Attributes:
  Strings:
    "source.ip": "172.0.0.1"
    "source.port": "9080"
    "source.name": "productpage"
    "source.uid": "001"
    "source.namespace": "default"
  StringMaps:
    "source.labels": 
      KeyValue:
        "disc": "ssd"
    "target.labels":
      KeyValue:
        "management": "backend"
Quotas:
- Name: "quota-name"
  Amount: 20
  BestEffort: true
`
func TestMixerClient(t *testing.T) {
	var checkRequest = pb.CheckRequest{}
	checkRequest03 = strings.ToLower(checkRequest03)
	err := yaml.Unmarshal([]byte(checkRequest03),&checkRequest)
	if err != nil {
		t.Errorf("%v" , err)
	}
	checkRequestjson, _ := json.MarshalIndent(checkRequest,"","   ")
	fmt.Printf("CheckeRequest: \n %s \n",checkRequestjson)
	checkResponse, err := MixerClient(checkRequest)

	checkResponsejson, _ := json.MarshalIndent(checkResponse,"","   ")
	fmt.Printf("CheckeResponse: \n %s \n",checkResponsejson)
}
