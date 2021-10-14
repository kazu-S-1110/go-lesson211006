package main

import "fmt"

// 型スイッチ

func anything(a interface{}) {
	// fmt.Println(a)
	// switch文を使って処理を変えてみる
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa") //interfaceは全てのものを受け入れられるが演算などは不可
	anything(1)

	// 型アサーションを使用して演算をさせることが可能
	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 4)
	// ちなみにこのxをfloat64でアサーションしてみるとruntimeerrorになる
	// 回避方法
	f, isFloat := x.(float64)
	fmt.Println(f, isFloat) //xをfloat64に変換することはできないがruntimeerrorを回避することは可能

	// 型のif
	if x == nil {
		fmt.Println("none")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("I don't know")
	}

	// 型のswitch(ifよりこちらが推奨)
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	// 値を受け取ることも可能
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println("I don't know'")
	}

}
