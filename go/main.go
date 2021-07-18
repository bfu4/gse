package main

import (
    "C"
    "unsafe"
    "test/mytls"
    "http"
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
func SendRequest(transport interface{}, creq *C.Request) unsafe.Pointer {
	methodString := C.GoString(*creq.method)
	site := C.GoString(*creq.url)
	body := C.GoString(*creq.body)
	asserted := transport.(*http.Transport)
	client := http.Client{Transport: &asserted}
	rdr := strings.NewReader(body)
	req, err := http.NewRequest(methodString, site, rdr)
	if err != nil {
		return unsafe.Pointer(nil)
	}

	// handle headers
	setHeaders(req, *creq.headers)

	// todo: (@bfu4) set the data

	// do the request
	resp, err := client.Do(req)
	// todo: (@bfu4) handle response
	return unsafe.Pointer(&resp)
}

// setHeaders set an array of headers
func setHeaders(req http.Request, headers []C.Header) {
	for k, v := range headers {
		setHeader(req, k)
	}
}

// setHeader set a single header
func setHeader(req http.Request, header C.Header) {
	key := C.GoString(header.key)
	value := C.GoString(header.value)
	req.Header.Set(key, value)
}

// ignored
func main() { }
