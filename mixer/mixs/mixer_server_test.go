// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package mixs

import (
	pb "demo_mesh/common-protos/mixer"
	"encoding/json"
	"fmt"
	"go-yaml/yaml"

	"strings"
	"testing"
)


func TestMixerServer_Check(t *testing.T) {
	var checkResponse = pb.CheckResponse{}
	checkResponse01 = strings.ToLower(checkResponse01)
	fmt.Printf("%s \n",checkResponse01)
	err := yaml.Unmarshal([]byte(checkResponse01),&checkResponse)
	if err != nil {
		t.Errorf("%v" , err)
	}

	checkResponsejson, _ := json.MarshalIndent(checkResponse,"","   ")
	fmt.Printf("%s \n",checkResponsejson)
}

func Test_Check(t *testing.T) {
	StartServer()

}