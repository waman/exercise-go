// 【練習問題 1.3】
// 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、
// 実行時間の差を計測しなさい（1.6節は time パッケージの一部を説明して
// いますし、11.4節では体系的に性能評価を行うためのベンチマークテストの書き方
// を説明しています）。
//
// 【実行方法】
// > cd ch1\exercise3
// > go test -bench=.
//
// 【結果】
// ・非効率な可能性のあるバージョン : 28324 ns/op
// ・strings.Join : 3959 ns/op
// 7倍以上の差
package exercise3

import (
	"testing"
	"strings"
)

// 連結を行う文字列のスライス。
// 練習問題 1.3 の時点では出てきてませんが。
func getSampleSliceOfString() []string {
	var arg []string
	for i := 0; i < 100; i++ {
		arg = append(arg, string(i))
	}
	return arg
}

func Benchmark非効率な可能性のあるバージョン(b *testing.B){
	args := getSampleSliceOfString()
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range args {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkStringsJoin(b *testing.B){
	args := getSampleSliceOfString()
	for i:= 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}