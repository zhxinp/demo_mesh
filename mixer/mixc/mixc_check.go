package mixc

import (
	"context"
	"golang.org/x/time/rate"
)

func Check() {
	var rateLimiter = rate.NewLimiter(1,1)
	rateLimiter.Wait(context.Background())
}
