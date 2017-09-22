// 【練習問題 1.3】
// 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、
// 実行時間の差を計測しなさい（1.6節は time パッケージの一部を説明して
// いますし、11.4節では体系的に性能評価を行うためのベンチマークテストの書き方
// を説明しています）。
//
// 【実行方法】（コマンド・プロンプト）
// > go get github.com/waman/exercise-go/ch1/ex1_3
// > cd %GOPATH%/ch1/ex1_3
// > go test -bench=.
//
// 【結果例】
// ・非効率な可能性のあるバージョン : 43172 ns/op
// ・strings.Join : 4157 ns/op
package ex1_3

import (
	"strconv"
	"strings"
	"testing"
)

// 連結を行う文字列のスライス。
// 練習問題 1.3 の時点では出てきてませんが。
var strSlice []string

// 文字列のスライスの初期化
// init() メソッド、strconv.Itoa() メソッドも
// 練習問題 1.3 の時点で出てきませんが。
func init() {
	for i := 0; i < 100; i++ {
		strSlice = append(strSlice, strconv.Itoa(i))
	}
}

func Benchmark非効率な可能性のあるバージョン(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range strSlice {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strSlice, " ")
	}
}
