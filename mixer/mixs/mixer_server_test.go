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