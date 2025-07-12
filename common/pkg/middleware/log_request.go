package middleware

import (
	"time"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func UseJSONLogFormat() {
	log.SetFormatter(&JSONFormatter{
		Program: "auth-service",
		Env:     "development",
	})

	log.SetLevel(log.DebugLevel)
}

var timeStampFormat = "2006-01-02T15:04:05.000000Z07:00"

type JSONFormatter struct {
	Program string
	Env     string
}

func (f *JSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["time"] = entry.Time.UTC().Format(timeStampFormat)
	data["msg"] = entry.Message
	data["level"] = strings.ToUpper(entry.Level.String())
	data["program"] = f.Program
	data["env"] = f.Env

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func GetClientIP(c *gin.Context) string {
	requester := c.Request.Header.Get("X-Forwarded-For")

	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}

	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if exists {
		return userID.(string)
	}
	return ""
}

func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

type LogRequestMiddleware struct{}

func (l LogRequestMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := GetDurationInMillseconds(start)

		entry := log.WithFields(log.Fields{
			"client_ip":   GetClientIP(c),
			"duration":    duration,
			"method":      c.Request.Method,
			"path":        c.Request.RequestURI,
			"status":      c.Writer.Status(),
			"user_id":     GetUserID(c),
			"referrer":    c.Request.Referer(),
			"request_id":  c.Writer.Header().Get("Request-Id"),
			"api_version": "1",
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
