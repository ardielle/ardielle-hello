//
// This file generated by rdl 1.4.12
//

package main

import (
	"net/http"

	hello "../.."
)

func main() {
	endpoint := "localhost:4080"
	url := "http://" + endpoint + "/hello"
	impl := new(hello.HelloImpl)
	handler := hello.Init(impl, url, impl)
	http.ListenAndServe(endpoint, handler)
}