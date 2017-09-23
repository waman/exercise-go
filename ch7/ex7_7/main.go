// 【練習問題 7.7】
// 20.0 のデフォルト値は°を含んでいないのに、ヘルプメッセージが°を含んでいる
// 理由を説明しなさい。
package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("摂氏%g度", c)  // 出力する文字列を変更
}

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

// ヘルプメッセージを出力するには以下を実行します：
//
//   > go run ./ch7/tempflag/main.go -help
//
// Celsius 型に対するメソッド String() を変更すればデフォルト値として
// 表示される値も変更されるので、CelsiusFlag に指定したデフォルト値 20.0
// は Celcius 型に変換され、ヘルプメッセージはその String() メソッドを
// 呼び出しているのが原因。
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
