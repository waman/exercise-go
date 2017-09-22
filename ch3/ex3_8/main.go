// 【練習問題 3.8】
// 高倍率の水準でフラクタルをレンダリングするには高い算術精度が
// 求められます。 complex64, complex128, big.Float, big.Rat
// の四つの異なる数値の表現を使って同じフラクタルを実装しなさい。
// （最後の二つの型は math/big パッケージにあります。 Float は
// 任意精度ですが有界精度の浮動小数点数を使っています。 Rat は
// 非有界精度の有理数を使っています。）性能とメモリ使用量に関して
// どのような比較結果になりますか。どの倍率の水準になるとレンダリング
// の結果が視覚的にわかるようになりますか。
package main

import (
	// 接頭辞なしで書けるようにするため
	"fmt"
)

func main() {
	fmt.Println("同ディレクトリ内の次のいずれかを実行してください。")
	fmt.Println("  mainC64.go")
	fmt.Println("  mainC128.go")
	fmt.Println("  mainFloat.go")
	fmt.Println("  mainRat.go")
}
