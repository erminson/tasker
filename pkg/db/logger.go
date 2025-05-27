package database

import (
	"fmt"
	"log/slog"
	"os"
)

type gooseLogger struct {
	*slog.Logger
}

func (l *gooseLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *gooseLogger) Printf(format string, v ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, v...))
}
