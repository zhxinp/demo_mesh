package webhook

type ValidationWebhook struct {
	ApiVersion string
	Metadata struct {
		Name string
	}
	Webhooks []struct {
		Name string
		ClientConfig struct {
			Service struct {
				Name string
				Namespace string
				Path string
			}
		}
		Rules [] struct {
			Operations [] string
			Resources [] string
		}
	}
}