package configration

type Rule struct {
	ApiVersion string
	Metadata *struct {
		Name string
		Namespace string
	}
	Spec struct {
		Match string
		Actions *[]struct {
			Handler string
			Instances *[]string
		}

	}
}
