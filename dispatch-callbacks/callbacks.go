package main

//#cgo darwin LDFLAGS: -F/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/System/Library/Frameworks/
//#cgo darwin LDFLAGS: -L/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib
//#include "callbacks.h"
//
// extern void invoke_callback_one(void*);
// extern void invoke_callback_two(void*, int);
// extern void invoke_callback_three(void*, int, char*);
import "C"
import (
	"log"
	"runtime/cgo"
	"unsafe"
)

var logObject = false

type Callbacks struct {
	One   func()
	Two   func(x int)
	Three func(x int, s string)
}

func InvokeCallbacks(cb *Callbacks) {
	ccallbacks := C.callbacks_unit{}

	ccallbacks.step_one_cb = C.step_one_cb_func(C.invoke_callback_one)
	ccallbacks.step_two_cb = C.step_two_cb_func(C.invoke_callback_two)
	ccallbacks.step_three_cb = C.step_three_cb_func(C.invoke_callback_three)

	handle := cgo.NewHandle(cb)

	C.invoke_callbacks(&ccallbacks, unsafe.Pointer(&handle))
}

//export invoke_callback_one
func invoke_callback_one(data unsafe.Pointer) {
	h := *(*cgo.Handle)(data)
	callbacks := h.Value().(*Callbacks)
	if logObject {
		log.Printf("%#v", callbacks)
	}

	if callbacks.One != nil {
		callbacks.One()
	}
}

//export invoke_callback_two
func invoke_callback_two(data unsafe.Pointer, x C.int) {
	h := *(*cgo.Handle)(data)
	callbacks := h.Value().(*Callbacks)
	if logObject {
		log.Printf("%#v", callbacks)
	}

	if callbacks.Two != nil {
		callbacks.Two(int(x))
	}
}

//export invoke_callback_three
func invoke_callback_three(data unsafe.Pointer, cx C.int, cs *C.char) {
	h := *(*cgo.Handle)(data)
	callbacks := h.Value().(*Callbacks)
	if logObject {
		log.Printf("%#v", callbacks)
	}

	if callbacks.Three != nil {
		callbacks.Three(int(cx), C.GoString(cs))
	}
}
