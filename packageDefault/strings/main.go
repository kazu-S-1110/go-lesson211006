package main

import (
	"fmt"
	"strings"
)

// strings

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C", "D"}, ",")
	s2 := strings.Join([]string{"A", "B", "C", "D"}, "")
	fmt.Println(s1, s2) //A,B,C,D ABCD

	// 文字列に含まれる部分文字列を検索する
	i1 := strings.Index("ABCD", "A")
	i2 := strings.Index("ABCD", "D")
	fmt.Println(i1, i2) //0 3

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3) //5

	i4 := strings.IndexAny("ABC", "ABC")
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5) //0 2

	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2) //true true

	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABCDE", "Bk")
	fmt.Println(b3, b4) //true true

	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7) //2 7
}
