package hippocampus

import (
	"log"
	"os"
)

func buildLoggerPrefix(name string) string {
	return "[" + name + "] "
}

func ChildLogger(parent *log.Logger, name string) *log.Logger {
	return log.New(parent.Writer(), buildLoggerPrefix(name), parent.Flags())
}

func NewLogger(name string) *log.Logger {
	return log.New(os.Stderr, buildLoggerPrefix(name), 0)
}
