// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	. "github.com/waman/exercise-go/ch2/tempconv"  // 接頭辞を省略
	"fmt"
	"flag"
)

// *celsiusFlag は flag.Value インターフェースを満足します。
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag は、指定された名前、デフォルト値、使い方を持つ Celsius フラグ
// を定義しており、そのフラグ変数のアドレスを返します。
// フラグ引数は度数と単位です。たとえば、"100C"です。
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

// 実行例：
//
//   > go build ./ch7/tempflag
//   > tempflag
//   > tempflag -temp -18C
//   > tempflag -temp 212°F
//
func main(){
	flag.Parse()
	fmt.Println(*temp)
}