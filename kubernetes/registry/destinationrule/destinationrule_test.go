package destinationrule

import (
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"strings"
	"testing"
)

var testDesinationrule01 = `
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews
spec:
  host: reviews
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
  - name: v3
    labels:
      version: v3
`

func TestDestinationrule(t *testing.T) {
	var testDestinationrule = Destinationrule{}
	testDesinationrule01 = strings.ToLower(testDesinationrule01)
	err := yaml.Unmarshal([]byte(testDesinationrule01),&testDestinationrule)
	if err != nil {
		t.Errorf("%v" , err)
	}
	testDestinationruleJson, _ := json.MarshalIndent(testDestinationrule,"","   ")
	fmt.Printf("testVirtualServiceJson: \n %s \n",testDestinationruleJson)
}