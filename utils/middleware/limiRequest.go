package middleware

import (
	"github.com/gin-gonic/gin"
	"goBackend-demo/gen/response"
	"sync"
	"time"
)

// RateLimiter 限流器
type RateLimiter struct {
	limiterMap map[string]*rateLimitInfo
	mutex      sync.Mutex
}

// rateLimitInfo 限流信息
type rateLimitInfo struct {
	count     int           // 计数
	firstSeen time.Time     // 第一次出现
	limit     int           // 限制
	TimeLimit time.Duration // 时间限制
}

// NewRateLimiter 创建限流器
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiterMap: make(map[string]*rateLimitInfo),
	}
}

// LimitRequest 限制请求
func (rl *RateLimiter) LimitRequest(limit int, TimeLimit time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		rl.mutex.Lock()
		defer rl.mutex.Unlock()
		info, exists := rl.limiterMap[ip]
		now := time.Now()
		if !exists || now.Sub(info.firstSeen) >= TimeLimit {
			info = &rateLimitInfo{
				count:     1,
				firstSeen: now,
				limit:     limit,
				TimeLimit: TimeLimit,
			}
			rl.limiterMap[ip] = info
		} else {
			if info.count >= limit {
				c.JSON(429, response.New("请求次数过多,请稍后再试", nil))
				c.Abort()
				return
			}
			info.count++
		}
		c.Next()
	}
}
