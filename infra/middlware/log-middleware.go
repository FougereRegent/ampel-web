package middlware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LogMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Before request
		route := ctx.Request.URL.Path
		method := ctx.Request.Method
		clientIp := ctx.ClientIP()

		logger.Info().
			Str("method", method).
			Str("route", route).
			Str("client ip", clientIp).
			Send()

		ctx.Next()
		//After request
		for _, err := range ctx.Errors {
			logger.Error().
				Err(err).
				Send()
		}
	}
}
