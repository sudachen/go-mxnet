package mx

import (
	"github.com/sudachen/go-mxnet/mx/internal"
)

const Version = 10401

func GpuCount() int {
	return internal.GpuCount
}

type Context int

func Cpu() Context {
	return Context(1)
}

func Gpu(devNo int) Context {
	return Context(2 + devNo*1000)
}

func (c Context) Int() int {
	return int(c)
}

func (c Context) DevType() int {
	return c.Int() % 1000
}

func (c Context) DevNo() int {
	return c.Int() / 1000
}
