package main

import (
	"fmt"
	"os"

	"github.com/ardielle/ardielle-hello/go"
)

func main() {
	endpoint := "localhost:4080"
	url := "http://" + endpoint + "/v1/hello"
	client := hello.NewClient(url, nil)
	name := os.Getenv("USER")
	greeting, err := client.GetGreeting(name)
	if err != nil {
		fmt.Println("***", err)
	} else {
		fmt.Println(greeting.Message)
	}
}
