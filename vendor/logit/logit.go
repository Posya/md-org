package logit

import (
	"log"
	"os"
)

var trace, info *log.Logger

func init() {
	f := os.Stdout
	trace = log.New(f, "T: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(f, "I: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info puts message to some log with INFO severity
func Info(s string) error {
	info.Println(s)
	return nil
}

func Trace(s string) error {
	trace.Println(s)
	return nil
}
