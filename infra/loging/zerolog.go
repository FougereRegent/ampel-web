package loging

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func New() zerolog.Logger {

	console := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFormatUnix,
		NoColor:    false,
	}

	console.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	console.FormatMessage = func(i any) string {
		if i == nil {
			i = ""
		}
		return fmt.Sprintf("%s", i)
	}
	console.FormatFieldName = func(i any) string {
		return fmt.Sprintf("%s:", i)
	}
	console.FormatTimestamp = func(i any) string {
		return fmt.Sprintf("%s", i)
	}

	logger := zerolog.New(console).
		With().
		Timestamp().
		Logger()

	return logger
}
