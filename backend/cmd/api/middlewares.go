package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *App) setMiddlewares() {
	a.router.Use(DefaultLogger())
	a.router.Use(gin.Recovery())
}

// DefaultLogger returns a gin middleware that logs requests using zerolog.
func DefaultLogger() gin.HandlerFunc {
	log.Info().Msg("Loading logger middleware...")
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
