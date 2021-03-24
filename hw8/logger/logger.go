// Package logger implements functions and structs to logging INFO ERROR WARNING messges to log file and to stdout
//
// The InitLoggers func opens(creates) log file and creates several loggers
//
//
//
package logger

import (
	"io"
	"log"
	"os"
)

var (
	// логгеры multiwrite std out and to file (az tee)
	WarningLogger  *log.Logger
	InfoLogger     *log.Logger
	ErrorLogger    *log.Logger
	// only file
	WarningFileLogger  *log.Logger
	InfoFileLogger     *log.Logger
	ErrorFileLogger    *log.Logger
)
// InitLoggers - начальная инициализация - лог файлы/ логгеры
// File - log file
func InitLoggers(File string) error {
	file, err := os.OpenFile(File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("error creating log file %v",err)
		return err
	}

	mw := io.MultiWriter(os.Stdout, file)
	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.LUTC)
	WarningLogger = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	InfoFileLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.LUTC)
	WarningFileLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	ErrorFileLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC|log.Lmicroseconds)

	return nil
}