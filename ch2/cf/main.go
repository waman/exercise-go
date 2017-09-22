// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// cf は、その数値引数を摂氏と華氏へ変換します。
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/waman/exercise-go/ch2/tempconv"
	// プロジェクトのルートディレクトリを GOPATH に設定しておけば
	// ch2/tempconv でインポートできると思うが・・・
)

func main() {
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
