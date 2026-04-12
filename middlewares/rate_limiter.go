package middlewares

import (
	"encoding/json"

	tollbooth "github.com/didip/tollbooth/v7"

	tollbooth_gin "github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}
type RateLimitConfig struct {
	Limit  float64
	Status string
	Body   string
}
func RateLimit(cfg RateLimitConfig) gin.HandlerFunc {
	msg := Message{
		Status: cfg.Status,
		Body:   cfg.Body,
	}
	jsonMsg, _ := json.Marshal(msg)

	limiter := tollbooth.NewLimiter(cfg.Limit, nil)
	limiter.SetMessageContentType("application/json")
	limiter.SetMessage(string(jsonMsg))

	return tollbooth_gin.LimitHandler(limiter)
}
