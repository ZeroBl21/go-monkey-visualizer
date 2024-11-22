package vm

import (
	"github.com/ZeroBl21/go-monkey-visualizer/internal/code"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
