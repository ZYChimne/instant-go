package api

import (
	"context"
	"time"
)

var ctx = context.Background()

const pageSize int64 = 10
const redisExpireTime time.Duration = 0 // 0 means no expire, ONLY FOR DEBUG
