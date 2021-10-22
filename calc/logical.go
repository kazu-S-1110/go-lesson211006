package main

import "fmt"

// 論理演算子

func main() {
	fmt.Println(true && false == true)  //false
	fmt.Println(true && false == false) //true

	fmt.Println(true || false)         //true
	fmt.Println(true || false == true) //true

	fmt.Println(!true)  //false
	fmt.Println(!false) //true
}
