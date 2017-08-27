// 【練習問題 2.3】
// 単一の式の代わりにループを使うように PopCount を書き直しなさい。
// 二つのバージョンの性能を比較しなさい。（11.4 節で異なる実装の
// 性能を体系的に比較する方法を説明しています。）
//
// 【注意】
// PopCount の各バージョンの実装は main.go に書いてます。
//
// 【実行方法】（コマンド・プロンプト）
// > go get github.com/waman/exercise-go/ch2/ex2_3
// > cd %GOPATH%/ch2/ex2_3
// > go test -bench=.
//
// 【結果例】
// ・単一の式を使ったバージョン : 0.87 ns/op
// ・ループを使ったバージョン : 113 ns/op
package ex2_3

import (
	"testing"
	"math/rand"
)

func Test単一の式を使ったバージョンとループを使ったバージョンの結果が等しい(t *testing.T){
	for i := 0; i < 100; i++ {
		x := rand.Uint64()
		if y, z := PopCount(x), PopCountByLoop(x); y != z {
			t.Errorf("結果が違います: %s: %s != %s", x, y, z)
		}
	}
}

var x = rand.Uint64()

func Benchmark単一の式を使ったバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func Benchmarkループを使ったバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCountByLoop(x)
	}
}