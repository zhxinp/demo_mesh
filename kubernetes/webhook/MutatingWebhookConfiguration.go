package webhook

type MutatingWebhook struct {
	ApiVersion string
	Metadata struct {
		Name string
		Labels map[string]string
	}
	Webhooks []struct {
		Name         string
		ClientConfig struct {
			Service struct {
				Name      string
				Namespace string
				Path      string
			}
		}
		Rules [] struct {
			Operations [] string
			Resources  [] string
		}
		NamespaceSelector struct {

			MatchLabels map[string]string
		}
	}
}


