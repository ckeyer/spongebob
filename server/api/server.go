package api

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	API_PREFIX = "/api"
)

func Serve(addr string) error {
	g := gin.New()
	g.Use(Logger(), Recovery())

	return g.Run(addr)
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method

		ctx.Next()

		if strings.HasPrefix(path, API_PREFIX) {
			end := time.Now()
			latency := end.Sub(start)

			statusCode := ctx.Writer.Status()
			comment := ctx.Errors.ByType(gin.ErrorTypePrivate).String()

			log.Debugf("%s %s %13v \"%s %s\" %v %s", clientIP, end.Format("2006/01/02-15:04:05"), latency, method, path, statusCode, comment)
		}
	}
}
