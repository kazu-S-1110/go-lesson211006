// testはgoの標準パッケージがある
// testファイルはテストしたいファイルと同一階層に作る
// ファイル名は(テストしたいファイル名)_test.goとすること

package main

import "testing"

var Debug bool = false //これはdebugをするかしないかの分岐

//テストする関数名はTest〜から始める
func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("Skip") //Debugがfalseなら下のコードが実行される
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}

	//テストの実行コマンドはgo test
	// go test -v とすると詳細に見れる
	// 出力結果
	// ok      _develop/go-lesson211006/test        0.308s

	// 失敗させてみた例
	// 	--- FAIL: TestIsOne (0.00s)
	//     forTest_test.go:20: 0 != 1
	// FAIL
	// exit status 1
	// FAIL    _/develop/go-lesson211006/test        0.336s
}
