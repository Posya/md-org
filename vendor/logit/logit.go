package logit

import (
	"log"
	"os"
)

var trace, info *log.Logger

func init() {
	trace = log.New(os.Stderr, "T: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stderr, "I: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Log puts message to some log with INFO severity
func Info(s string) error {
	info.Println(s)
	return nil
}

func Trace(s string) error {
	trace.Println(s)
	return nil
}
