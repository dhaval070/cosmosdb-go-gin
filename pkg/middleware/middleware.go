package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func Telemetry(client appinsights.TelemetryClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		c.Set("requestID", reqID)
		t := time.Now()
		c.Next()

		url := c.Request.URL.String()
		duration := time.Since(t)
		status := fmt.Sprintf("%d", c.Writer.Status())

		client.TrackRequest(c.Request.Method, url, duration, status)
	}
}
