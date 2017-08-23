// 【練習問題 2.1】
// 絶対温度 (Kelvin scale) で温度を処理するために tempconv に
// 型、定数、関数を追加しなさい。0K は -273.15°C であり、1K の
// 差と 1°C の差は同じ大きさです。
package main

import (
  . "github.com/waman/exercise-go/ch2/exercise1/tempconv1"
	  // tempconv1. と書くのが面倒なので、修飾名がいらないようにインポートしてます。
	"fmt"
)

func main(){
	fmt.Printf("絶対零度は %gK\n", AbsoluteZero)
	fmt.Printf("水の融点は %gK\n", Freezing)
	fmt.Printf("水の沸点は %gK\n", Boiling)
	fmt.Println()

	fmt.Printf("絶対零度は %g°C\n", KToC(AbsoluteZero))
	fmt.Printf("水の融点は %g°C\n", KToC(Freezing))
	fmt.Printf("水の沸点は %g°C\n", KToC(Boiling))
	fmt.Println()

	fmt.Printf("絶対零度は %g°F\n", KToF(AbsoluteZero))
	fmt.Printf("水の融点は %g°F\n", KToF(Freezing))
	fmt.Printf("水の沸点は %g°F\n", KToF(Boiling))
}
