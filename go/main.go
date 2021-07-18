package main

import (
    // #cgo CFLAGS: -I../c/include
    // #include <http-signatures.h>
    "C"
    "unsafe"
    "reflect"
    "gsego/mytls"
    http "net/http"
    "strings"
)

//export CreateJA3Transport
func CreateJA3Transport(ja3 *C.char) unsafe.Pointer {
	gostr := C.GoString(ja3)
	transport, _ := mytls.NewTransport(gostr)
	addr := unsafe.Pointer(transport)
	return addr
}

//export SendRequest
func SendRequest(transport unsafe.Pointer, creq *C.Request) unsafe.Pointer {
	methodString := C.GoString(creq.method)
	site := C.GoString(creq.url)
	body := C.GoString(creq.body)
	asserted := (*http.Transport) (rawType(transport))
	casted := **(**http.RoundTripper) (rawType(asserted))
	client := http.Client{Transport: casted}
	rdr := strings.NewReader(body)
	req, err := http.NewRequest(methodString, site, rdr)

	if err != nil {
		return unsafe.Pointer(nil)
	}

	// handle headers
	hdrs := **(**[]C.Header) (rawType((*creq).headers))
	setHeaders(req, hdrs)

	// todo: (@bfu4) set the data

	// do the request
	resp, err := client.Do(req)
	// todo: (@bfu4) handle response
	return unsafe.Pointer(&resp)
}

func rawType(value interface{}) unsafe.Pointer {
	gil := reflect.ValueOf(value)
	ptr := gil.Pointer()
	addrof := struct {
		addr uintptr
	}{ptr}
	veryRaw := unsafe.Pointer(&addrof)
	return veryRaw
}

// setHeaders set an array of headers
func setHeaders(req *http.Request, headers []C.Header) {
	for _, v := range headers {
		setHeader(req, v)
	}
}

// setHeader set a single header
func setHeader(req *http.Request, header C.Header) {
	key := C.GoString(header.name)
	value := C.GoString(header.value)
	req.Header.Set(key, value)
}

// ignored
func main() { }
