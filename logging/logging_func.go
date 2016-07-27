package logging

import (
	"io"
	"log"
)

var (
	initRun = false
)

// Init initialises the log writers for qilbot.
func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	initRun = true

	const (
		FLAGS = log.Ldate | log.Ltime | log.Lshortfile
	)

	Trace = log.New(traceHandle, "TRACE: ", FLAGS)
	Info = log.New(infoHandle, "INFO: ", FLAGS)
	Warning = log.New(warningHandle, "WARNING: ", FLAGS)
	Error = log.New(errorHandle, "ERROR: ", FLAGS)
}
