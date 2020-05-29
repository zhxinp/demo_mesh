package adapter

import (
	rpc "googleapis/google/rpc"
	"time"
)

//调用handler之后的返回值
type QuotaResult struct {
	//操作的结果状态
	Status rpc.Status

	//quota的有效期时间间隔，如果是0，就表示永不过期
	ValidDuration time.Duration

	//quota返回的总数量，有可能比请求的数量少
	Amount int64
}
