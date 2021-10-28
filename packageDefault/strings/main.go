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

	// 文字列を繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4) //ABCABCABCABC

	// 文字列の置換
	s5 := strings.Replace("AAAAA", "A", "ggd", 1)
	s6 := strings.Replace("AAAAA", "A", "ggd", -1) //-1を指定した場合、全箇所を置換
	fmt.Println(s5, s6)                            //ggdAAAA ggdggdggdggdggd

	// 文字列を分割する
	s7 := strings.Split("A,B,C,D,E,F,G", ",")  //[A B C D E F G]
	s8 := strings.SplitAfter("A,B,C,D,E", ",") //[A, B, C, D, E]
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)  //[A B,C,D,E]
	s10 := strings.SplitN("A,B,C,D,E", ",", 4) //[A B C D,E]
	fmt.Println(s7, s8, s9, s10)

	// 大文字小文字の変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")
	s13 := strings.ToUpper("abcdka")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14) //abc e ABCDKA E

	// 文字列から空白を取り除く(前後のみ)
	s15 := strings.TrimSpace("     - ---- abads -----    ")
	fmt.Println(s15)
	// 全角
	s16 := strings.TrimSpace("     　　　Abcgad   っv 　　")
	fmt.Println(s16)

	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1]) //スライスとして取り出すのでIndex指定が可能になる
}
