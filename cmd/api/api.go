package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	listenPort  int = 8080
	rootCommand     = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}
)

func init() {
	rootCommand.Flags().IntVarP(&listenPort, "listen-port", "p", 8080, "this flags is used to set a listen port")
}

func main() {
	rootCommand.Execute()
	connectionString := fmt.Sprintf("0.0.0.0:%d", listenPort)
	router := gin.New()

	configAssetsAndTemplates(router)
	configRoute(router)

	router.Run(connectionString)
}

func configAssetsAndTemplates(router *gin.Engine) {
	router.Static("/static", "./web/assets")
	router.LoadHTMLFiles("web/templates/index.html")
}

func configRoute(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "alive",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
