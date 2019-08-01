package tests

import (
	"fmt"
	"github.com/sudachen/go-mxnet/mx"
	"gotest.tools/assert"
	"testing"
)

func Test_Version(t *testing.T) {
	assert.Equal(t, mx.Version, 10401)
}

func Test_Array1(t *testing.T) {
	a := mx.Array(mx.Int32, mx.Dim(100))
	assert.NilError(t, a.Err())
	assert.Assert(t, a != nil)
	assert.Assert(t, a.Dtype().String() == "Int32")
	assert.Assert(t, a.Dim().String() == "(100)")
}

func Test_Array2(t *testing.T) {
	a := mx.Array(mx.Float32, mx.Dim(100, 10))
	assert.NilError(t, a.Err())
	assert.Assert(t, a != nil)
	assert.Assert(t, a.Dtype().String() == "Float32")
	assert.Assert(t, a.Dim().String() == "(100,10)")
}

func Test_Array3(t *testing.T) {
	a := mx.Array(mx.Uint8, mx.Dim(100, 10, 1))
	assert.NilError(t, a.Err())
	assert.Assert(t, a.Dtype().String() == "Uint8")
	assert.Assert(t, a.Dim().String() == "(100,10,1)")
}

func Test_Array4(t *testing.T) {
	a := mx.Array(mx.Int64, mx.Dim(100000000000, 100000000, 1000000000))
	assert.ErrorContains(t, a.Err(), "failed to create array")
}

func Test_Array5(t *testing.T) {
	var a *mx.NDArray
	a = mx.Array(mx.Int64, mx.Dim())
	assert.ErrorContains(t, a.Err(), "bad dimension")
	a = mx.Array(mx.Int64, mx.Dim(-1, 3))
	assert.ErrorContains(t, a.Err(), "bad dimension")
	a = mx.Array(mx.Int64, mx.Dim(1, 3, 10, 100, 2))
	assert.ErrorContains(t, a.Err(), "bad dimension")
}

func Test_Random_Uniform(t *testing.T) {
	a := mx.Array(mx.Float32, mx.Dim(1, 3)).Uniform(0, 1)
	assert.NilError(t, a.Err())
	fmt.Println(a)
}
