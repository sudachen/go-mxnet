package tests

import (
	"github.com/sudachen/go-mxnet/mxnet"
	"gotest.tools/assert"
	"testing"
)

func Test_Version(t *testing.T) {
	assert.Equal(t, mxnet.Version, 10401)
}

func Test_Array1(t *testing.T) {
	a, err := mxnet.Cpu().Array(mxnet.Int32, mxnet.Dim(100))
	assert.NilError(t, err)
	assert.Assert(t, a != nil)
	assert.Assert(t, a.Type().String() == "Int32")
	assert.Assert(t, a.Dim().String() == "(100)")
}

func Test_Array2(t *testing.T) {
	a, err := mxnet.Cpu().Array(mxnet.Float32, mxnet.Dim(100, 10))
	assert.NilError(t, err)
	assert.Assert(t, a != nil)
	assert.Assert(t, a.Type().String() == "Float32")
	assert.Assert(t, a.Dim().String() == "(100,10)")
}

func Test_Array3(t *testing.T) {
	a, err := mxnet.Cpu().Array(mxnet.Uint8, mxnet.Dim(100, 10, 1))
	assert.NilError(t, err)
	assert.Assert(t, a != nil)
	assert.Assert(t, a.Type().String() == "Uint8")
	assert.Assert(t, a.Dim().String() == "(100,10,1)")
}

func Test_Array4(t *testing.T) {
	_, err := mxnet.Cpu().Array(mxnet.Int64, mxnet.Dim(100000000000, 100000000, 1000000000))
	assert.ErrorContains(t, err, "failed to create array")
}

func Test_Array5(t *testing.T) {
	var err error
	_, err = mxnet.Cpu().Array(mxnet.Int64, mxnet.Dim())
	assert.ErrorContains(t, err, "bad dimension")
	_, err = mxnet.Cpu().Array(mxnet.Int64, mxnet.Dim(-1, 3))
	assert.ErrorContains(t, err, "bad dimension")
	_, err = mxnet.Cpu().Array(mxnet.Int64, mxnet.Dim(1, 3, 10, 100))
	assert.ErrorContains(t, err, "bad dimension")
}
