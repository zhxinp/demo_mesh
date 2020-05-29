// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com

package mixc


import (
	"context"
	"golang.org/x/time/rate"
)

func Check() {
	var rateLimiter = rate.NewLimiter(1,1)
	rateLimiter.Wait(context.Background())
}
