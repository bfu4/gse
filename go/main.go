package main

import (
    "C"
    "unsafe"
    "test/mytls"
)

//export CreateJA3Transport
func CreateJA3Transport(ja3 *C.char) C.long {
	gostr := C.GoString(ja3)
	transport, _ := mytls.NewTransport(gostr)
	addr := unsafe.Pointer(transport)
	goLong := uintptr(addr)
	hailMary := C.long(goLong)
	return hailMary
}

func main() { }
