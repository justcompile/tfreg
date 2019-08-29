package logging

type logLevel string

var (
	logLevelDebug = logLevel("DEBUG")
	logLevelError = logLevel("ERROR")
	logLevelInfo  = logLevel("INFO")
)
