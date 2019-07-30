package mx

import (
	"fmt"
)

const (
	DimRow = 0
	DimColumn = 1
	DimDepth = 2
)

// do not change this constant
// code can assume exactly this value
const MaxDimensionCount = 4

type Dimension struct {
	Shape [MaxDimensionCount]int
	Len   int
}

func (dim Dimension) String() string {
	s := "(%d,%d,%d,%d"[0:(3+(dim.Len-1)*3)] + ")"
	q := ([]interface{}{dim.Shape[0], dim.Shape[1], dim.Shape[2], dim.Shape[3]})[:dim.Len]
	return fmt.Sprintf(s, q...)
}

func (dim Dimension) Good() bool {
	if dim.Len <= 0 || dim.Len > MaxDimensionCount {
		return false
	}
	for _, v := range dim.Shape[:dim.Len] {
		if v <= 0 {
			return false
		}
	}
	return true
}

func (dim Dimension) SizeOf(dt Dtype) int {
	r := dt.Size()
	for i:=0; i<dim.Len; i++ {
		r *= dim.Shape[i]
	}
	return r
}

func Dim(a ...int) Dimension {
	var dim Dimension
	if q := len(a); q > 0 && q <= MaxDimensionCount {
		dim.Len = q
		for i, v := range a {
			dim.Shape[i] = v
		}
	}
	return dim
}

