package logging

import (
	"log/slog"
	"os"
)

var (
	InfoLogger  *slog.Logger
	ErrorLogger *slog.Logger
	DbLogger    *slog.Logger
)

func InitLoggers() error {

	infoFile, err := os.OpenFile("../../logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	errorFile, err := os.OpenFile("../../logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	dbFile, err := os.OpenFile("../../logs/db.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	InfoLogger  = slog.New(slog.NewTextHandler(infoFile, nil))
	ErrorLogger = slog.New(slog.NewTextHandler(errorFile, nil))
	DbLogger  	= slog.New(slog.NewTextHandler(dbFile, nil))
	
	return nil
}
