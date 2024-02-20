package metrics

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PromotheusMiddleware() gin.HandlerFunc {
	metrics := NewMetrics(NamespaceHTTP)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		metrics.IncRequestCounter(c.Request.Method, c.FullPath(), fmt.Sprint(c.Writer.Status()))
		metrics.ObserveRequestDuration(c.Request.Method, c.FullPath(), fmt.Sprint(c.Writer.Status()), duration)
	}
}
