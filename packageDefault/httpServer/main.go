package main

import (
	"fmt"
	"net/http"
)

// net/http
// server side

type MyHandler struct{}

// これの問題点はURLのどこにアクセスしても同じ画面になる、なのでParseFileを使う
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	// http.HandleFunc("/top",top)
	http.ListenAndServe(":8080", &MyHandler{})
}
