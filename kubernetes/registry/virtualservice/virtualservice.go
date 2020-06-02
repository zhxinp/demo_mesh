package virtualservice

type VirtualService struct {
	ApiVersion string
	Kind string
	Metadata struct {
		Name string
	}
	Spec struct {
		Hosts [] string
		Http [] struct {
			Match [] struct {
				Headers struct {
					Enduser struct {
						Exact string
				    }
				}
			}
			Fault struct {
				Delay struct {
					Percentage struct {
						Value float32
					}
					FixedDelay string
				}
				Abort struct {
					Percentage struct {
						Value float32
					}
					HttpStatus string
				}
			}
			Route [] struct {
				Destination struct {
					Host string
					Subset string
				}
				Weight int
			}

		}
	}
}