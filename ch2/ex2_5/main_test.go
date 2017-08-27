// 【練習問題 2.5】
// 式 x&(x-1) は x で1が設定されている最下位ビットをクリアします。
// この事実を使ってビット数を数える PopCount のバージョンを作成し、
// その性能を評価しなさい。
//
// 【注意】
// PopCount の各バージョンの実装は main.go に書いてます。
//
// 【実行方法】
// > cd ch2/exercise5
// > go test -bench=.
//
// 【結果例】
// ・通常のバージョン : 0.86 ns/op
// ・ビットクリアによるバージョン : 305 ns/op
package exercise3

import (
	"testing"
	"math/rand"
)

func Testビットクリアによるバージョンが正しい結果を与える(t *testing.T){
	for i := 0; i < 100; i++ {
		x := rand.Uint64()
		if y, z := PopCount(x), PopCountWithBitClear(x); y != z {
			t.Errorf("結果が違います: %s: %s != %s", x, y, z)
		}
	}
}

var x = rand.Uint64()

func Benchmark通常のバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func Benchmarkビットクリアによるバージョン(b *testing.B){
	for i := 0; i < b.N; i++ {
		PopCountWithBitClear(x)
	}
}