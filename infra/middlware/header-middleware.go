package middlware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	applicationContent string = "Content-Type"
	jsonApplication    string = "application/json"
)

func HeaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		switch ctx.Request.Method {
		case "POST":
			if !applicationContentIsSet(ctx, jsonApplication) {
				return
			}
		case "PUT":
			if !applicationContentIsSet(ctx, jsonApplication) {
				return
			}
		case "PATCH":
			if !applicationContentIsSet(ctx, jsonApplication) {
				return
			}
		default:
			break
		}
		ctx.Next()
	}
}

func applicationContentIsSet(ctx *gin.Context, value string) bool {
	appContent := ctx.Request.Header.Get(applicationContent)
	result := appContent == value
	if !result {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
	}
	return result
}
