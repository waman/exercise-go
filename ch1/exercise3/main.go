// 【練習問題 1.3】
// 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、
// 実行時間の差を計測しなさい（1.6節は time パッケージの一部を説明して
// いますし、11.4節では体系的に性能評価を行うためのベンチマークテストの書き方
// を説明しています）。
package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

func main(){
	// 非効率な可能性のあるバージョン
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("非効率な可能性のあるバージョンの実行時間：%.5fs\n", secs)

	// strings.Join
	start = time.Now()
	strings.Join(os.Args, " ")
	secs = time.Since(start).Seconds()
	fmt.Printf("strings.Join バージョンの実行時間：%.5fs\n", secs)
}