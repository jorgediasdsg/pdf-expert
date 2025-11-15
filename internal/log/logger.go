package log

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init(env string) {
	if env == "prod" {
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
}
