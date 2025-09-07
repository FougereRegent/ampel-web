package main

import (
	"ampel-web/infra/loging"
	"ampel-web/infra/middlware"
	"ampel-web/internal/domain"
	"ampel-web/internal/usecases"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	listenPort  int = 8080
	rootCommand     = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}
	logger zerolog.Logger
)

func init() {
	rootCommand.Flags().IntVarP(&listenPort, "listen-port", "p", 8080, "this flags is used to set a listen port")
	logger = loging.New()
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

func colorMapping(color string) (domain.LedColor, error) {
	color = strings.ToLower(color)
	switch color {
	case "green":
		return domain.Green, nil
	case "red":
		return domain.Red, nil
	case "orange":
		return domain.Orange, nil
	default:
		return -1, errors.New("This color cannot exist")
	}
}

func configRoute(router *gin.Engine) {
	ledUseCases := usecases.New()
	api := router.Group("/api")
	api.Use(middlware.LogMiddleware(&logger), middlware.HeaderMiddleware(), middlware.ErrorMiddleware())
	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "alive",
		})
	})

	api.POST("/led/on", func(ctx *gin.Context) {
		var ledColor struct {
			Color string `json:"color"`
		}
		if err := ctx.BindJSON(&ledColor); err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		color, err := colorMapping(ledColor.Color)
		if color == -1 {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		evt := domain.EnableLedEvent{
			Color: color,
		}

		if err = ledUseCases.TreatLedEvent(&evt); err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Status(http.StatusOK)
	})

	api.POST("/led/off", func(ctx *gin.Context) {
		var ledColor struct {
			Color string `json:"color"`
		}
		if err := ctx.BindJSON(&ledColor); err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		color, err := colorMapping(ledColor.Color)
		if color == -1 {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		evt := domain.DisbaleLedEvent{
			Color: color,
		}

		if err = ledUseCases.TreatLedEvent(&evt); err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Status(http.StatusOK)
	})

	router.GET("reference", func(ctx *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.yml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Ampel-led API",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		ctx.Writer.Write([]byte(htmlContent))
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
