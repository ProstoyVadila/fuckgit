package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *App) setMiddlewares() {
	a.router.Use(loggerMiddleware())
	a.router.Use(CORSMiddleware())
	a.router.Use(gin.Recovery())
}

// loggerMiddleware returns a gin middleware that logs requests using zerolog.
func loggerMiddleware() gin.HandlerFunc {
	log.Info().Msg("Loading logger middleware")
	logger := &log.Logger
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		logParams := gin.LogFormatterParams{}
		logParams.TimeStamp = time.Now()
		logParams.Latency = logParams.TimeStamp.Sub(start)
		logParams.ClientIP = ctx.ClientIP()
		logParams.Method = ctx.Request.Method
		logParams.StatusCode = ctx.Writer.Status()
		logParams.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		logParams.BodySize = ctx.Writer.Size()
		if raw != "" {
			path += "?" + raw
		}
		logParams.Path = path

		var logEvent *zerolog.Event
		switch {
		case logParams.StatusCode >= 500:
			logEvent = logger.Error()
		case logParams.StatusCode >= 400:
			logEvent = logger.Warn()
		default:
			logEvent = logger.Info()
		}
		logEvent.
			Str("client_ip", logParams.ClientIP).
			Str("method", logParams.Method).
			Int("status_code", logParams.StatusCode).
			Int("body_size", logParams.BodySize).
			Str("path", logParams.Path).
			Str("latency", logParams.Latency.String()).
			Msg(logParams.ErrorMessage)
	}
}

// CORSMiddleware returns a gin middleware that sets CORS headers.
func CORSMiddleware() gin.HandlerFunc {
	log.Info().Msg("Loading CORS middleware")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
