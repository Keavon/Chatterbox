package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	// Ldebug is debug loglevel
	Ldebug = 0
	// Linfo is info loglevel.
	Linfo = 1
	// Lwarn is debug loglevel.
	Lwarn = 1
	// Lerror is debug loglevel.
	Lerror = 2
	// Mconsole is console output mode
	Mconsole = 0
	// Log format
	format = log.Ldate | log.Ltime | log.Lshortfile
)

var (
	// Debug Logger
	Debug *log.Logger
	// Info Logger
	Info *log.Logger
	// Warn Logger
	Warn *log.Logger
	// Error Logger
	Error *log.Logger
)

// New creates a new logger
func New(mode int, loglevel int) {
	Debug = log.New(ioutil.Discard, "", format)
	Info = log.New(ioutil.Discard, "", format)
	Warn = log.New(ioutil.Discard, "", format)
	Error = log.New(ioutil.Discard, "", format)

	if mode == Mconsole {
		if loglevel <= Ldebug {
			Debug = log.New(os.Stdout, "Debug: ", format)
			Debug.Print("Debug logging started...")
		}
		if loglevel <= Linfo {
			Info = log.New(os.Stdout, "Info: ", format)
			Info.Print("Info logging started...")
		}
		if loglevel <= Lwarn {
			Warn = log.New(os.Stderr, "Warn: ", format)
			Warn.Print("Warning logging started...")
		}
		if loglevel <= Lerror {
			Error = log.New(os.Stderr, "Error: ", format)
			Error.Print("Error logging started...")
		}
	}
}

// StringToFlag converts names for log settings into flags
func StringToFlag(logOutput, logLevel string) (int, int, error) {
	lo := Mconsole
	if logOutput != "console" {
		return -1, -1, fmt.Errorf("Invalid logOutput: %s\n", logOutput)
	}

	switch logLevel {
	case "debug":
		return lo, Ldebug, nil
	case "info":
		return lo, Linfo, nil
	case "warn":
		return lo, Lwarn, nil
	case "error":
		return lo, Lerror, nil
	default:
		return -1, -1, fmt.Errorf("Invalid logLevel: %s\n", logLevel)
	}
}
