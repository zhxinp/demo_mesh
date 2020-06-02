package webhook

import (
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"strings"
	"testing"
)

var ValidationWebhookConfiguration01 = `
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: istio-galley
webhooks:
  - name: pilot.validation.istio.io
    clientConfig:
      service:
        name: istio-galley
        namespace: istio-system
        path: "/admitpilot"
    rules:
      - operations:
        - CREATE
        - UPDATE
        resources:
        - httpapispecs
        - httpapispecbindings
        - quotaspecs
        - quotaspecbindings
      - operations:
        - CREATE
        - UPDATE
        resources:
        - destinationrules
        - envoyfilters
        - gateways
        - serviceentries
        - sidecars
        - virtualservices
  - name: mixer.validation.istio.io
    clientConfig:
      service:
        name: istio-galley
        namespace: istio-system
        path: "/admitmixer"
      caBundle: ""
    rules:
      - operations:
        - CREATE
        - UPDATE
        resources:
        - rules
        - attributemanifests
        - adapters
        - handlers
        - instances
        - templates
`


func Test_ValidationWebHook(t *testing.T) {
	var testValidationWebHook = ValidationWebhook{}
	ValidationWebhookConfiguration01 = strings.ToLower(ValidationWebhookConfiguration01)
	err := yaml.Unmarshal([]byte(ValidationWebhookConfiguration01),&testValidationWebHook)
	if err != nil {
		t.Errorf("%v" , err)
	}
	testValidationWebHookJson, _ := json.MarshalIndent(testValidationWebHook,"","   ")
	fmt.Printf("testValidationWebHookJson: \n %s \n",testValidationWebHookJson)
}