package main

// #cgo CFLAGS: -O0
// #cgo LDFLAGS: -L. -lcstring_concat
// char* string_concat(char*, char*);
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := C.CString("Hello")
	b := C.CString("World")
	ret := (C.string_concat(a, b))
	goRet := C.GoString(ret)
	fmt.Println(goRet)
	C.free(unsafe.Pointer(a))
	C.free(unsafe.Pointer(b))
	C.free(unsafe.Pointer(ret))
	//fmt.Println(C.fibo(C.int(10)))
}
