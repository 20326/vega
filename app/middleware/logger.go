package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

const headerXRequestID = "X-Request-ID"

// Logger is the logrus logger handler
func LoggerWithRequestID(logger logrus.FieldLogger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		reqID := c.Request.Header.Get(headerXRequestID)
		if reqID == "" {
			reqID = ksuid.New().String()
			c.Header(headerXRequestID, reqID)
		}
		start := time.Now()
		c.Next()
		end := time.Now()
		statusCode := c.Writer.Status()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"time":       end.Format(time.RFC3339),
			"latency":    latency,
			"reqID":      reqID,
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%d] %s %s [%d] (%dms)", statusCode, c.Request.Method, path, dataLength, latency)
			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
