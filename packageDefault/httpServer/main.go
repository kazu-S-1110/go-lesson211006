package main

import (
	"html/template"
	"log"
	"net/http"
)

// net/http
// server side

// type MyHandler struct{}

// // これの問題点はURLのどこにアクセスしても同じ画面になる、なのでParseFileを使う
// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111!") //tmpl.htmlの{{}}の.を書き換えているらしい

}

func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}
