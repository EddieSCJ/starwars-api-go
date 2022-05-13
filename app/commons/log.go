package commons

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroLogger() zerolog.Logger {
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "message"
	zerolog.CallerFieldName = "file"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().
		Caller().
		Stack().
		Str("app", "Star Wars API").
		Str("env", GetEnvironment()).
		Logger()

	if strings.ToLower(GetEnvironment()) == "dev" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().
			Caller().
			Stack().
			Str("app", "Star Wars API").
			Str("env", GetEnvironment()).
			Logger()
	}

	return logger
}
