package log

import (
	"fmt"
	"log/slog"
	"os"
)

func SetupLogger(logPath string) *slog.Logger {
	var log slog.Logger
	logData, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(logPath)
		panic("could not open log file" + err.Error())
	}
	log = *slog.New(slog.NewTextHandler(logData, &slog.HandlerOptions{AddSource: false, Level: slog.LevelInfo}))
	return &log
}
