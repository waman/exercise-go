// 【練習問題 2.1】
// 絶対温度 (Kelvin scale) で温度を処理するために tempconv に
// 型、定数、関数を追加しなさい。0K は -273.15°C であり、1K の
// 差と 1°C の差は同じ大きさです。
package tempconv

import (
	"fmt"
)

// このファイルのコードは【練習問題 7.6】で使うためにパッケージとして公開しています。
// コード例は main_test.go にあります。　このコードは
//
//   > go test ./ch2/ex2_1/tempconv
//
// によってテストできます。

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Kelvin = 0
	Freezing     Kelvin = 273.15
	Boiling      Kelvin = 373.15
)

func (t Kelvin) String() string     { return fmt.Sprintf("%gK", t) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// KToC は絶対温度を摂氏へ変換します
func KToC(t Kelvin) Celsius { return Celsius(float64(t) - float64(Freezing)) }

// CToK は摂氏を絶対温度へ変換します
func CToK(c Celsius) Kelvin { return Kelvin(float64(c) + float64(Freezing)) }

// CToF は摂氏を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32)) * 5 / 9 }

// KToF は絶対温度を華氏へ変換します
func KToF(t Kelvin) Fahrenheit { return CToF(KToC(t)) }

// FToK は華氏を絶対温度へ変換します
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }