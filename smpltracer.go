package smpltrace

import "io"

// simple tracing interface

type Smpltracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Smpltracer {
	return nil
}
