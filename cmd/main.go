package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
