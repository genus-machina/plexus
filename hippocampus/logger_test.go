package hippocampus

import (
	"os"
	"testing"
)

func TestRootLogger(t *testing.T) {
	logger := NewLogger("test")
	assertPrefix(t, logger, "[test] ")
	assertFlags(t, logger, 0)
	assertWriter(t, logger, os.Stderr)
}

func TestChildLogger(t *testing.T) {
	parent := NewLogger("parent")
	child := ChildLogger(parent, "child")
	assertPrefix(t, child, "[child] ")
	assertFlags(t, child, parent.Flags())
	assertWriter(t, child, parent.Writer())
}
