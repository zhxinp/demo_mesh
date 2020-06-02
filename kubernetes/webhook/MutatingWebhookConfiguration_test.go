package webhook

import (
	"encoding/json"
	"fmt"
	"go-yaml/yaml"
	"strings"
	"testing"
)

var MutatingWebhookConfiguration01 = `
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: istio-sidecar-injector
  labels:
    app: sidecarInjectorWebhook
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      service:
        name: istio-sidecar-injector
        namespace: istio-system
        path: "/inject"
    rules:
      - operations: [ "CREATE" ]
        resources: ["pods"]
    namespaceSelector:
      matchLabels:
        istio-injection: enabled
`

func TestMutatingWebhookConfiguration(t *testing.T) {
	var testMutatingWebHook = MutatingWebhook{}
	MutatingWebhookConfiguration01 = strings.ToLower(MutatingWebhookConfiguration01)
	err := yaml.Unmarshal([]byte(MutatingWebhookConfiguration01),&testMutatingWebHook)
	if err != nil {
		t.Errorf("%v" , err)
	}
	testMutatingWebHookJson, _ := json.MarshalIndent(testMutatingWebHook,"","   ")
	fmt.Printf("testMutatingWebHookJson: \n %s \n",testMutatingWebHookJson)
}