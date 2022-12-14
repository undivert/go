package main

import (
	"io"
	"os"
	"log"
)

func main() {
	// Call the writeLog function to write an info log message.
	writeLog("info", "Writing an info log message.")

	// Call the writeLog function to write a warning log message.
	writeLog("warning", "Writing a warning log message.")

	// Call the writeLog function to write an error log message.
	writeLog("error", "Writing an error log message.")
}

// writeLog writes a log message of the specified type to a file in a designated "logs" directory.
// The log level must be one of "info", "warning", or "error".
func writeLog(level string, message string) {
	// Create a logs directory if it doesn't already exist.
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Open a log file in the logs directory.
	logFile, err := os.OpenFile("logs/myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	// Create a logger for each log level.
	var logger *log.Logger
	switch level {
	case "info":
		logger = log.New(io.MultiWriter(logFile, os.Stdout), "INFO: ", log.LstdFlags)
	case "warning":
		logger = log.New(io.MultiWriter(logFile, os.Stdout), "WARNING: ", log.LstdFlags)
	case "error":
		logger = log.New(io.MultiWriter(logFile, os.Stdout), "ERROR: ", log.LstdFlags)
	default:
		log.Fatalf("invalid log level: %s", level)
	}

	// Write the log message.
	logger.Println(message)
}