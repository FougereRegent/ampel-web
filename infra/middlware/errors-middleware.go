package middlware

import (
	"ampel-web/internal/execptions"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {

	type errorMessage struct {
		ErrorMessage string `json:"error-message"`
	}
	return func(ctx *gin.Context) {
		//Before request
		ctx.Next()
		// After Request
		for _, err := range ctx.Errors {
			if errors.Is(err, execptions.BadRequestError) {
				ctx.JSON(http.StatusBadRequest, errorMessage{
					ErrorMessage: err.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, errorMessage{
					ErrorMessage: "Internal Server Error",
				})
				return
			}
		}
	}
}
