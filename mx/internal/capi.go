package internal

/*
#cgo CFLAGS: -I/opt/mxnet/include
#cgo LDFLAGS: -L/opt/mxnet/lib -lmxnet -Wl,-rpath=/opt/mxnet/lib
#include <mxnet/c_api.h>

static int imperative_invoke_out1(AtomicSymbolCreator ent, NDArrayHandle out, int ano, const char **keys, const char **vals) {
	NDArrayHandle* out1[1] = {&out};
	int nout = 1;
	int err = MXImperativeInvoke(ent, 0, NULL, &nout, &out1[0], ano, keys, vals);
	return err;
}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

var GpuCount int = 0
var LibVersion = 0
var mxkeys [KeyNoKey]*C.char
var mxentry [OpNoOp]C.AtomicSymbolCreator

type NDArrayHandle C.NDArrayHandle

func init() {

	var v C.int
	C.MXGetVersion(&v)
	LibVersion = int(v)

	var c C.int
	C.MXGetGPUCount(&c)
	GpuCount = int(c)

	for i := KeyEmpty + 1; i < KeyNoKey; i++ {
		mxkeys[i] = C.CString(i.Value())
	}

	var ascv *C.AtomicSymbolCreator
	var ascn C.uint

	if e := C.MXSymbolListAtomicSymbolCreators(&ascn, &ascv); e != 0 {
		panic("failed to gather symbols from mxnet")
	}

	m := map[string]MxnetOp{}
	for op := OpEmpty + 1; op < OpNoOp; op++ {
		m[op.Value()] = op
	}

	for i := uintptr(0); i < uintptr(ascn); i++ {
		a := *(*C.AtomicSymbolCreator)(unsafe.Pointer(uintptr(unsafe.Pointer(ascv)) + i*unsafe.Sizeof(*ascv)))
		var n *C.char
		if e := C.MXSymbolGetAtomicSymbolName(a, &n); e != 0 {
			panic(fmt.Sprintf("failed to gather name for symbol %x", a))
		}
		if ent, ok := m[C.GoString(n)]; ok {
			mxentry[ent] = a
			fmt.Println(C.GoString(n), a)
		}
	}
}

const maxArgsCount = 16

func fillargs(keys []*C.char, vals []*C.char, ap []interface{}) int {
	i := 0
	for len(ap) != 0 && i < maxArgsCount/2 {
		keys[i] = mxkeys[ap[0].(MxnetKey)]
		vals[i] = C.CString(fmt.Sprint(ap[1]))
		i++
		ap = ap[2:]
	}
	return i
}

func ImperativeInvokeInplace1(op MxnetOp, h NDArrayHandle, a ...interface{}) error {
	if h == nil {
		return fmt.Errorf("uninitialized or broken array")
	}

	var keys [maxArgsCount]*C.char
	var vals [maxArgsCount]*C.char
	ano := C.int(fillargs(keys[:], vals[:], a))
	if ent := mxentry[op]; ent != nil {
		if e := C.imperative_invoke_out1(ent, C.NDArrayHandle(h), ano, &keys[0], &vals[0]); e != 0 {
			return fmt.Errorf("maxnet api error: %v", op.Value())
		}
	} else {
		return fmt.Errorf("unresolved API entry %v", op.Value())
	}
	return nil
}

/*func ImperativeInvokeOut1(op MxnetOp, out *NDArrayHandle, in []NDArrayHandle, a... interface{}) error {
	var keys [16]*C.char
	var vals [16]*C.char
	ap := a
	i := 0
	for len(ap) != 0 {
		keys[i] = mxkeys[ap[0].(MxnetKey)]
		vals[i] = C.CString(fmt.Sprint(ap[1]))
		i++
		ap = ap[2:]
	}
	var cin [2]C.NDArrayHandle
	var cout [2]*C.NDArrayHandle
	cout[0] = (*C.NDArrayHandle)(unsafe.Pointer(&cout[1]))
	nout := C.int(1)
	if ent := mxentry[op]; ent != nil {
		e := C.MXImperativeInvoke(
			ent,
			C.int(len(in)),&cin[0],
			&nout,&cout[0],
			C.int(i),&keys[0],&vals[0])
		if e != 0 {
			return fmt.Errorf("maxnet api error: %v", op.Value())
		}
	} else {
		return fmt.Errorf("unresolved API entry %v", op.Value())
	}
	return nil
}
*/

func NewNDArrayHandle(devType int, devNo int, dtype int, shape [4]int, slen int) (NDArrayHandle, int) {
	var a C.NDArrayHandle
	s := [4]C.uint{C.uint(shape[0]), C.uint(shape[1]), C.uint(shape[2]), C.uint(shape[3])}
	e := C.MXNDArrayCreateEx(&s[0], C.uint(slen), C.int(devType), C.int(devNo), 0, C.int(dtype), &a)
	return NDArrayHandle(a), int(e)
}

func ReleaseNDArrayHandle(handle NDArrayHandle) {
	if handle != nil {
		C.MXNDArrayFree(C.NDArrayHandle(handle))
	}
}

