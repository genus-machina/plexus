package hippocampus

import (
	"os"
	"testing"
)

func TestRootLogger(t *testing.T) {
	logger := NewLogger("test")

	if prefix := logger.Prefix(); prefix != "[test] " {
		t.Errorf("expected prefix '[test] ' but got '%s'", prefix)
	}

	if flags := logger.Flags(); flags != 0 {
		t.Errorf("expected flags 0 but got %d", flags)
	}

	if writer := logger.Writer(); writer != os.Stderr {
		t.Error("writer was not stderr")
	}
}

func TestChildLogger(t *testing.T) {
	parent := NewLogger("parent")
	child := ChildLogger(parent, "child")

	if prefix := child.Prefix(); prefix != "[child] " {
		t.Errorf("expected prefix '[child] ' but got '%s'", prefix)
	}

	if flags := child.Flags(); flags != parent.Flags() {
		t.Errorf("expected flags %d but got %d", parent.Flags(), flags)
	}

	if writer := child.Writer(); writer != parent.Writer() {
		t.Error("writers do not match")
	}
}
