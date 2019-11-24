package hippocampus

import (
	"io"
	"log"
	"testing"
)

func assertPrefix(t *testing.T, logger *log.Logger, expected string) {
	if prefix := logger.Prefix(); prefix != expected {
		t.Errorf("expected prefix '%s' but got '%s'", expected, prefix)
	}
}

func assertFlags(t *testing.T, logger *log.Logger, expected int) {
	if flags := logger.Flags(); flags != expected {
		t.Errorf("expected flags %d but got %d", expected, flags)
	}
}

func assertWriter(t *testing.T, logger *log.Logger, expected io.Writer) {
	if writer := logger.Writer(); writer != expected {
		t.Error("writers do not match")
	}
}
