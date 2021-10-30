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

	// 正規表現のフラグ
	// フラグ一覧
	//  i 大文字小文字を区別しない
	//  m マルチラインモード（^,&が文頭、文末に加えて行頭、行末にマッチ）
	//  s .が\nにマッチ
	//  U 最小マッチへの変換

	re3 := regexp.MustCompile(`(?i)abc`) //(?XXX)の部分でフラグを指定（複数もおk）
	// (?i-ms)の場合、iは有効にしてmsを無効化する
	match = re3.MatchString("ABC")
	fmt.Println(match)

	// 幅を持たない正規表現のパターン
	// パターン一覧
	// ^ 文頭
	// $ 文末
	// \A 文頭
	// \z 文末
	// \b ASCIIによるワード協会
	// \B 非ASCIIによるワード協会

	re4 := regexp.MustCompile(`^ABC`) //文頭にABCがあるか
	match = re4.MatchString("ABCkdfjasldkjflas")
	fmt.Println(match)
	match = re4.MatchString("kdfjasldkjflas")
	fmt.Println(match)

	// 繰り返しを表す正規表現
	// 繰り返しのパターン

	// x*　0回以上繰り返すx（最大マッチ
	// x+　1回以上繰り返すx（最大マッチ
	// x?　0回以上1回以下繰り返すx
	// x{n,m} n回以上m回以下繰り返すx（最大マッチ
	// x{n,}　n回以上繰り返すx（最大マッチ
	// x{n}　n回繰り返す（最大マッチ
	// x*?　0回以上繰り返すマッチ（最小マッチ
	// x+?　1回以上繰り返すマッチ（最小マッチ
	// x??　0回以上1回以下繰り返すマッチ（0回優先
	// x{n,m}?　n回以上m回以下繰り返すx（最小マッチ
	// x{n,}?　n回以上繰り返すx（最小マッチ
	// x{n}?　n回繰り返すx（最小マッチ

	re6 := regexp.MustCompile("a+b*")
	fmt.Println("ab", re6.MatchString("ab"))                     //true
	fmt.Println("a", re6.MatchString("a"))                       //true
	fmt.Println("aaabaaabbbbb", re6.MatchString("aaabaaabbbbb")) //true
	fmt.Println("b", re6.MatchString("b"))                       //false
}
