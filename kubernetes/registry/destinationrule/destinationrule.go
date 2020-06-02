package destinationrule

type Destinationrule struct {
	ApiVersion string
	Kind string
	Metadata struct {
		Name string

	}
	Spec struct {
		Host string
		Subsets []struct {
			Name string
			Labels struct {
				Version string
			}
		}
	}
}
