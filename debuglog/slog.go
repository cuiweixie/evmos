package debuglog

import (
	"golang.org/x/exp/slog"
	"os"
)

var log *slog.Logger

func init() {
	file, err := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log = slog.New(slog.NewTextHandler(file, nil))
}

func GetLogger() *slog.Logger {
	return log
}
