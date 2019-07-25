package mxnet

// #cgo CFLAGS: -I${SRCDIR}/../include
// #cgo LDFLAGS: -lmxnet -L${SRCDIR}/../lib
// #include <mxnet/c_api.h>
import "C"

const Version = 10401

func init() {
}
