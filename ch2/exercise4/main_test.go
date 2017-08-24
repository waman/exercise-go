// 【練習問題 2.4】
// 引数をビットシフトしながら最下位ビットの検査を64回繰り返すことで
// ビット数を数える PopCount のバージョンを作成しなさい。
// テーブル参照を行うバージョンと性能を比較しなさい。
//
// 【注意】
// PopCount の各バージョンのは main.go に書いてます。
//
// 【実行方法】
// > cd ch2/exercise4
// > go test -bench=.
//
// 【結果】
// ・テーブル参照を行うバージョン : 0.87 ns/op
// ・テーブル参照を行わないバージョン : 295 ns/op
package exercise3

import (
	"testing"
	"math/rand"
)

func Testテーブル参照を行ったバージョンと行っていないバージョンの結果が等しい(t *testing.T){
	for i := 0; i < 100; i++ {
		x := rand.Uint64()
		if PopCount(x) != PopCountWithoutTable(x) {
			t.Errorf("結果が違います: %s: %s != %s", x, PopCount(x), PopCountWithoutTable(x))
		}
	}
}

var x = rand.Uint64()

func Benchmarkテーブル参照を行うバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func Benchmarkテーブル参照を行わないバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCountWithoutTable(x)
	}
}