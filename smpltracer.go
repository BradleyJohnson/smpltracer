package smpltrace

import (
	"fmt"
	"io"
)

// simple tracing interface

type Smpltracer interface {
	Trace(...interface{})
}

type smpltracer struct {
	out io.Writer
}

func (t *smpltracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func New(w io.Writer) Smpltracer {
	return &smpltracer{out: w}
}
