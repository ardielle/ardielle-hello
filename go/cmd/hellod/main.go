package main

import (
	"net/http"

	"github.com/ardielle/ardielle-hello/go"
)

func main() {
	endpoint := "localhost:4080"
	url := "http://" + endpoint + "/v1/hello"
	impl := new(hello.HelloImpl)
	handler := hello.Init(impl, url, impl)
	http.ListenAndServe(endpoint, handler)
}
