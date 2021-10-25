package main

import (
	alib "alib/mypkg" //なんでエラー吐いているのに実行できるんだよ、、、
	"fmt"
)

//test

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(3))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
