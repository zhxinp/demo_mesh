package virtualservice

import (
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"strings"
	"testing"
)

var testVertualService01 = `
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews
spec:
  hosts:
  - reviews
  http:
  - match:
    - headers:
         enduser:
           exact: lucy
    route:
    - destination:
        host: reviews
        subset: v2
      weight: 50
    - destination:
        host: reviews
        subset: v3
      weight: 50
  - match:
      - headers:
          enduser:
            exact: jason
    fault:
      delay:
        percentage:
          value: 100.0
        fixedDelay: 2s
    route:
      - destination:
          host: reviews
          subset: v2
  - match:
      - headers:
          enduser:
            exact: judy
    fault:
      abort:
        percentage:
          value: 100.0
        httpStatus: 500
    route:
      - destination:
          host: reviews
          subset: v1
  - route:
    - destination:
        host: reviews
        subset: v1
`
func TestVertualService(t *testing.T) {
	var testVirtualService = VirtualService{}
	testVertualService01 = strings.ToLower(testVertualService01)
	err := yaml.Unmarshal([]byte(testVertualService01),&testVirtualService)
	if err != nil {
		t.Errorf("%v" , err)
	}
	testVirtualServiceJson, _ := json.MarshalIndent(testVirtualService,"","   ")
	fmt.Printf("testVirtualServiceJson: \n %s \n",testVirtualServiceJson)
}