package mxnet

// #cgo CFLAGS: -I${SRCDIR}/../include
// #include <mxnet/c_api.h>
import "C"

import (
	"fmt"
	"runtime"
)

type Dtype int

const (
	Float32 Dtype = 0
	Float64 Dtype = 1
	Float16 Dtype = 2
	Uint8   Dtype = 3
	Int32   Dtype = 4
	Int8    Dtype = 5
	Int64   Dtype = 6
)

func (tp Dtype) String() string {
	switch tp {
	case Float32:
		return "Float32"
	case Float64:
		return "Float64"
	case Float16:
		return "Float16"
	case Uint8:
		return "Uint8"
	case Int32:
		return "Int32"
	case Int8:
		return "Int8"
	case Int64:
		return "Int64"
	default:
		panic("bad type")
	}
}

type Dimension struct {
	Shape [3]int
	Len   int
}

func Dim(a ...int) Dimension {
	var dim Dimension
	if q := len(a); q > 0 && q <= 3 {
		dim.Len = q
		for i, v := range a {
			dim.Shape[i] = v
		}
	}
	return dim
}

func (dim Dimension) String() string {
	s := "(%d,%d,%d"[0:(3+(dim.Len-1)*3)] + ")"
	q := ([]interface{}{dim.Shape[0], dim.Shape[1], dim.Shape[2]})[:dim.Len]
	return fmt.Sprintf(s, q...)
}

func (dim Dimension) Good() bool {
	if dim.Len <= 0 || dim.Len > 3 {
		return false
	}
	for _, v := range dim.Shape[:dim.Len] {
		if v <= 0 {
			return false
		}
	}
	return true
}

type Array struct {
	ctx    Context
	dim    Dimension
	dtype  Dtype
	handle C.NDArrayHandle
}

type Context struct {
	devType int
	devNo   int
}

func Cpu() Context {
	return Context{1, 0}
}

func Gpu(devNo int) Context {
	return Context{2, devNo}
}

func (c Context) Array(tp Dtype, d Dimension) (*Array, error) {
	if !d.Good() {
		return nil, fmt.Errorf("failed to create array %v%v: bad dimension", tp.String(), d.String())
	}
	a := &Array{ctx: c, dim: d, dtype: tp}
	s := [3]C.uint{C.uint(d.Shape[0]), C.uint(d.Shape[1]), C.uint(d.Shape[2])}
	e := C.MXNDArrayCreateEx(&s[0], C.uint(d.Len), C.int(c.devType), C.int(c.devNo), 0, C.int(tp), &a.handle)
	if e != 0 {
		return nil, fmt.Errorf("failed to create array %v%v: api error", tp.String(), d.String())
	}
	runtime.SetFinalizer(a, func(a *Array) { a.Release() })
	return a, nil
}

func (a *Array) Type() Dtype {
	return a.dtype
}

func (a *Array) Dim() Dimension {
	return a.dim
}

func (a *Array) Release() {
	if a.handle != nil {
		C.MXNDArrayFree(a.handle)
		a.handle = nil
	}
}

// idiomatic finalizer
func (a *Array) Close() error {
	a.Release()
	return nil
}
