package main

import (
	f "fmt" //パッケージ名の短縮（あまり推奨されない）
	foo "foo/mypkg"
)

func main() {
	f.Println(foo.Max)
	// fmt.Println(foo.min) //min not exported by package foo
	f.Println(foo.ReturnMin()) //1
}
