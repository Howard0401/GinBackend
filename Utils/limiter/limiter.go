package limiter

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//https://github.com/juju/ratelimit

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterInterface interface {
	Key(c *gin.Context) string //
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterInterface
}

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterInterface {
	l := Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return &MethodLimiter{
		Limiter: &l,
	}
}

//獲取RequestURI作為Key
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

//查找桶(Buckut)中是否存在Key Value
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}
	return l
}
