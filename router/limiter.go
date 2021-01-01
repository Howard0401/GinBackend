package router

import (
	"VueGin/Utils/limiter"
	"time"
)

var Limiter = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		//限制同時登入數
		Key:          "/login/auth",
		FillInterval: time.Second,
		Capacity:     100,
		Quantum:      100,
	},
)
