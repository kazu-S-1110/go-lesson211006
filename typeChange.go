package main

import (
	"fmt"
	"strconv"
)

// 型変換
func main() {
	var i int = 1
	//int => float
	fl64 := float64(i)
	fmt.Println(fl64)               //1
	fmt.Printf("fl64 = %T\n", fl64) //float64

	//float => int
	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2) //int

	//float int 浮動小数ありでも可能（小数は切り捨て
	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3) //int
	fmt.Println(i3)             //10

	// string => int ちょっと特殊
	var s string = "100"
	fmt.Printf("s = %T\n", s) //string
	i4, _ := strconv.Atoi(s)  //返り値は2つあるが、2つ目の受け取りを_とすることで受け取ったけど破棄している
	// i4, err := strconv.Atoi(s)  //本来の形
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(i)
	fmt.Printf("i4 = %T\n", i4) //int

	// int => string
	var i5 int = 200
	s2 := strconv.Itoa(i5)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2) //string

	// string => byte
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) //[72 101 108 108 111 32 87 111 114 108 100]

	// byte => string
	h2 := string(b)
	fmt.Println(h2) //Hello World
}
