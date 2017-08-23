// 【練習問題 2.1】
// 絶対温度 (Kelvin scale) で温度を処理するために tempconv に
// 型、定数、関数を追加しなさい。0K は -273.15°C であり、1K の
// 差と 1°C の差は同じ大きさです。
package main

import (
	"os"
	"strconv"
	"fmt"

	"github.com/waman/exercise-go/ch2/tempconv"
	// プロジェクトのルートディレクトリを GOPATH に設定しておけば
	// ch2/tempconv でインポートできると思うが・・・
)

func main(){
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
