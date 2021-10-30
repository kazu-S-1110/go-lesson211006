package main

import (
	"fmt"
	"regexp"
)

// regexp

func main() {
	// GOの正規表現の基本 regexp.matchString (大規模には向かない書き方)
	// match, _ := regexp.MatchString("A", "ABC")
	// fmt.Println(match) //true

	// Complie
	// re1, _ := regexp.Compile("A")
	// match = re1.MatchString("ABC")
	// fmt.Println(match) //true

	// mustCompile(エラーを返さずruntimeerrを出す。動的に使うならこっちがいいかも)
	// re2 := regexp.MustCompile("A")
	// match = re2.MatchString("ABC")
	// fmt.Println(match)

	// バックスラッシュの表現方法
	// regexp.MustCompile("\\d")
	// regexp.MustCompile(`\d`)

	// 正規表現のフラグ
	// フラグ一覧
	//  i 大文字小文字を区別しない
	//  m マルチラインモード（^,&が文頭、文末に加えて行頭、行末にマッチ）
	//  s .が\nにマッチ
	//  U 最小マッチへの変換

	// re3 := regexp.MustCompile(`(?i)abc`) //(?XXX)の部分でフラグを指定（複数もおk）
	// (?i-ms)の場合、iは有効にしてmsを無効化する
	// match = re3.MatchString("ABC")
	// fmt.Println(match)

	// 幅を持たない正規表現のパターン
	// パターン一覧
	// ^ 文頭
	// $ 文末
	// \A 文頭
	// \z 文末
	// \b ASCIIによるワード協会
	// \B 非ASCIIによるワード協会

	// re4 := regexp.MustCompile(`^ABC`) //文頭にABCがあるか
	// match = re4.MatchString("ABCkdfjasldkjflas")
	// fmt.Println(match)
	// match = re4.MatchString("kdfjasldkjflas")
	// fmt.Println(match)

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

	// re6 := regexp.MustCompile("a+b*")
	// fmt.Println("ab", re6.MatchString("ab"))                     //true
	// fmt.Println("a", re6.MatchString("a"))                       //true
	// fmt.Println("aaabaaabbbbb", re6.MatchString("aaabaaabbbbb")) //true
	// fmt.Println("b", re6.MatchString("b"))                       //false

	// 正規表現の文字クラス
	// re7 := regexp.MustCompile(`[XYZ]`)
	// fmt.Println("Y", re7.MatchString("Y"))
	// fmt.Println("AB", re7.MatchString("AB"))
	// fmt.Println("YZ", re7.MatchString("YZ"))

	// re8 := regexp.MustCompile(`^[0-9A-Za-z_]{4}$`)
	// fmt.Println("ABC", re8.MatchString("ABC"))
	// fmt.Println("aABC", re8.MatchString("aABC"))
	// fmt.Println("あ", re8.MatchString("あ"))

	// re9 := regexp.MustCompile("[^0-9A-Za-z_]") //[]括弧内先頭に^をつけるとそれ以外を指定する
	// fmt.Println("ABC", re9.MatchString("ABC")) //false
	// fmt.Println("(')", re9.MatchString("(')")) //true

	// re10 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	// fmt.Println(re10.MatchString("abcxyz"))
	// fmt.Println(re10.MatchString("ABCxyz"))
	// fmt.Println(re10.MatchString("abcXYZ"))
	// fmt.Println(re10.MatchString("ABCABC"))

	// 正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`)                      //\sは空白文字
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) //AAA,BBB,CCC

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("私は、Golangを学習する、プログラマーです。。よろしく、お願い。いたします。。", "")) //私はGolangを学習するプログラマーですよろしくお願いいたします

	// 正規表現による文字列の分割
	re3 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re3.Split("ASHVJBDK<LKJFABCXYZGLJA>KGDLSABCXYZ", -1)) //[ASHVJBDK<LKJF GLJA>KGDLS ]

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("alkfdalkdfla  adkj  glaksjdfla gasdlkfjalkgas   fgalsdkjfalskd", -1)) //[alkfdalkdfla adkj glaksjdfla gasdlkfjalkgas fgalsdkjfalskd]
}
