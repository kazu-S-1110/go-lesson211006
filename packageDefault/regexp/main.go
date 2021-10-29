package main

import (
	"fmt"
	"regexp"
)

// regexp

func main() {
	// GOの正規表現の基本 regexp.matchString (大規模には向かない書き方)
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match) //true

	// Complie
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match) //true

	// mustCompile(エラーを返さずruntimeerrを出す。動的に使うならこっちがいいかも)
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// バックスラッシュの表現方法
	// regexp.MustCompile("\\d")
	// regexp.MustCompile(`\d`)
}
