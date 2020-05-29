package main




import (
"context"
"fmt"
"time"

"golang.org/x/time/rate"
)

func main() {
    //b为初始化令牌个数，也是最大能存储的令牌个数，r是每秒钟放几个令牌进去。没执行一个动作就要消耗掉一个令牌
	l := rate.NewLimiter(1, 1)
	ctx := context.Background()
	start := time.Now()
	// 要处理二十个事件
	for i := 0; i < 20; i++ {
		l.WaitN(ctx,1)
		// dosomething
		fmt.Printf("i:%d \n", i)
	}
	fmt.Println(time.Since(start)) // 总执行时间
}