package mx

import (
	"fmt"
	"github.com/sudachen/go-mxnet/mx/internal"
	"runtime"
)

type NDArray struct {
	ctx	   Context
	dim    Dimension
	dtype  Dtype
	handle internal.NDArrayHandle
	err    error
}

func release(a *NDArray) { internal.ReleaseNDArrayHandle(a.handle) }

// idiomatic finalizer
func (a *NDArray) Close() error {
	release(a)
	return nil
}

func (a *NDArray) Release() {
	release(a)
}

func (a NDArray) Err() error {
	return a.err
}

func Array(tp Dtype, d Dimension) *NDArray {
	return Cpu().Array(tp, d)
}

func (c Context) Array(tp Dtype, d Dimension) *NDArray {
	if !d.Good() {
		return &NDArray{err: fmt.Errorf("failed to create array %v%v: bad dimension", tp.String(), d.String())}
	}
	a := &NDArray{ctx: c, dim: d, dtype: tp}
	if h,e := internal.NewNDArrayHandle(c.DevType(), c.DevNo(), int(tp), d.Shape, d.Len); e != 0 {
		return &NDArray{err: fmt.Errorf("failed to create array %v%v: api error", tp.String(), d.String())}
	} else {
		a.handle = h
		runtime.SetFinalizer(a, release)
	}
	return a
}

func (a *NDArray) Dtype() Dtype {
	return a.dtype
}

func (a *NDArray) Dim() Dimension {
	return a.dim
}

func (a *NDArray) Cast(dt Dtype) *NDArray {
	return nil
}

func (a *NDArray) Reshape(dim Dimension) *NDArray {
	return nil
}

func (a *NDArray) Data() []byte {
	return nil
}

type Variant struct {
	Value interface{}
}

func (v Variant) Float32() float32 {
	switch x := v.Value.(type) {
	case float32:
		return x
	case float64:
		return float32(x)
	case int64:
		return float32(x)
	}
	return 0
}

func (a *NDArray) Get(idx... int) Variant {
	return Variant{0}
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
func (a *NDArray) String() string {
	rs := [][]string{}
	lr := a.Len(DimRow)
	wr := min(lr,5)
	lc := a.Len(DimColumn)
	wc := min(lc,5)
	for row:=0; row<wr; row++ {
		rc := make([]string,wc)
		r := row
		if row == 2 && lr > 5 {
			for col := 0; col < wc; col++ {
				rc[col] = ".."
			}
			rs = append(rs,rc)
			continue
		} else if row > 3 && lr > 5 {
			r = lr - 5 + row
		}
		for col := 0; col < wc; col++ {
			if col < 2 || lc <= 5 {
				rc[col] = fmt.Sprintf("%v",a.Get(r, col).Value)
			} else if col == 2 {
				rc[col] = ".."
			} else {
				rc[col] = fmt.Sprintf("%v",a.Get(r, lc-5+col).Value)
			}
		}
		rs = append(rs,rc)
	}
	return fmt.Sprint(rs)
}

func (a *NDArray) Len(d int) int {
	if d < 0 || d >=3 {
		return 0
	}
	if a.dim.Len <= d { return 1 }
	return a.dim.Shape[d]
}

func (a *NDArray) Size() int {
	return a.dim.SizeOf(a.dtype)
}

