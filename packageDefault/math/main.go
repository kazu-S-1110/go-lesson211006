package main

import (
	"fmt"
	"math"
)

// math

func main() {
	// 数学的な定数
	// 円周率
	fmt.Println(math.Pi) //3.141592653589793

	// 2の平方根
	fmt.Println(math.Sqrt(2)) //1.4142135623730951
	fmt.Println(math.Sqrt(3)) //1.7320508075688772
	fmt.Println(math.Sqrt(5)) //2.23606797749979

	// 数値型に関する定数
	// float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)
	// float32で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	// float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)
	// float64で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64)
	// int ver
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100)) //100

	// 累乗を求める
	fmt.Println(math.Pow(0, 2)) //0
	fmt.Println(math.Pow(2, 4)) //16

	// 平方根、立方根
	fmt.Println(math.Sqrt(3))
	fmt.Println(math.Cbrt(8)) //2

	//最大値、最小値
	fmt.Println(math.Max(1, 3)) //3
	fmt.Println(math.Min(1, 3)) //1

	// 小数点以下の切り捨て、切り上げ

	// math.Trunc 数値の正負を問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.1415926535))  //3
	fmt.Println(math.Trunc(-3.1415926535)) //-3

	// math.Floor 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))  //3
	fmt.Println(math.Floor(-3.5)) //-4

	// math.Ceil 引数の数値より大きいもの最小の整数を返す
	fmt.Println(math.Ceil(3.5))  //4
	fmt.Println(math.Ceil(-3.5)) //-3

}
