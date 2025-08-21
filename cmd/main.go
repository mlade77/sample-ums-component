package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	err := http.ListenAndServe(":8555", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
