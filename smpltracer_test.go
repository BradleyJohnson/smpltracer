package smpltrace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	smpltracer := New(&buf)
	if smpltracer == nil {
		t.Error("Return from New() should not be nil.")
	} else {
		smpltracer.Trace("Hello world")
		if buf.String() != "Hello world\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}
