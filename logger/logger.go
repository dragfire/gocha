package logger

import (
	_ "io"
	"log"
	"os"
)

var (
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
