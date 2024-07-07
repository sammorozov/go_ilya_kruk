package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	// Using standard error for error logs and standard output for info and warning logs.
	errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.LstdFlags)
	infoLogger  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

// HaltOnErr logs an error and exits if the error is non-nil.
func HaltOnErr(err error, messages ...string) {
	if err == nil {
		return
	}

	message := "An error occurred"

	if len(messages) > 0 {
		message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))
	}
	errorLogger.Printf("%s: %v", message, err)
	os.Exit(1)
}

// Info logs an informational message.
func Info(message string) {
	infoLogger.Println(message)
}

// Warn logs a warning message along with an error if provided.
func Warn(err error, messages ...string) {
	if err != nil {
		message := "A warning occurred"
		if len(messages) > 0 {
			message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))
		}
		warnLogger.Printf("%s: %v", message, err)
	}
}
