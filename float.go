package main

import "fmt"

func main() {
	var fl64 float64 = 2.4

	fl := 3.2 //floatの暗黙的宣言では自動でfloat64になる、環境依存ではないので注意

	fmt.Println(fl64 + fl)
	fmt.Printf("%T\n", fl)

	// 特殊な値を保持する場合
	zero := 0.0
	pint := 1.0 / zero
	fmt.Println(pint) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN
}
