// 【練習問題 4.2】
// デフォルトで標準入力の SHA256 ハッシュを表示するプログラムを書きなさい。
// ただし、SHA384 ハッシュや SHA512 ハッシュを表示するコマンドラインの
// フラグもサポートしなさい。
package main

import (
	"fmt"
	"crypto/sha256"
	"os"
	"crypto/sha512"
)

// 入力文字列の SHA256 ハッシュを表示します。
// 第2引数に 384 もしくは 512 を指定すると、それぞれ SHA384, SHA512 ハッシュ
// を表示します。
func main(){
	var flag string = ""
	if len(os.Args) >=3 { flag = os.Args[2] }

	switch flag {
	case "384":
		fmt.Printf("SHA384: %x", sha512.Sum384([]byte(os.Args[1])))
	case "512":
		fmt.Printf("SHA512: %x", sha512.Sum512([]byte(os.Args[1])))
	default:
		fmt.Printf("SHA256: %x", sha256.Sum256([]byte(os.Args[1])))
	}
}
