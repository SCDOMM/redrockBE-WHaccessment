package utils

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func MiddleHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		path := string(c.Request.Path())
		method := string(c.Request.Method())

		c.Next(ctx)

		latency := time.Since(start)
		status := c.Response.StatusCode()
		log.Printf("[%s] %s %d %v", method, path, status, latency)
	}
}
