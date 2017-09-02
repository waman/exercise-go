// 【練習問題 4.1】
// 二つの SHA256 ハッシュで異なるビットの数を数える関数を
// 書きなさい（2.6.2節の PopCount 参照）。
package main

import (
	"crypto/sha256"
	"os"
	"fmt"
)

var pc [256]byte

func init(){
	for i := range pc { pc[i] = pc[i/2] + byte(i&1)	}
}

// 2.6.2節参照
func PopCount(x byte) int {
	return int(pc[x])
}

func PopCountDifference(x, y [32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += PopCount(x[i]^y[i])  // XOR (^) は異なるビットのみを1にする
	}
	return sum
}

func main(){
	if len(os.Args) != 3 {
		fmt.Println("2つの引数を指定してください。")
		return
	}

	var c1 [32]byte = sha256.Sum256([]byte(os.Args[1]))
	var c2 [32]byte = sha256.Sum256([]byte(os.Args[2]))

	fmt.Printf("SHA256 ハッシュの異なるビットの数： %d\n", PopCountDifference(c1, c2))
}
