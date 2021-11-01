package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// net/http
// client

func main() {
	// get method
	// res, _ := http.Get("http://example.com")
	// fmt.Println(res)
	// fmt.Println(res.StatusCode)             //200
	// fmt.Println(res.Proto)                  //HTTP/1.1
	// fmt.Println(res.Header["Date"])         //[Mon, 01 Nov 2021 08:06:45 GMT]
	// fmt.Println(res.Header["Content-Type"]) //[text/html; charset=UTF-8]

	// fmt.Println(res.Request.Method) //GET
	// fmt.Println(res.Request.URL)    //http://example.com

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	// post method
	vs := url.Values{}

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode()) //上で作ったものをencode

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body1, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body1))
}
