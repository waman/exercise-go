// 【練習問題 3.13】
// できるだけコンパクトに KB, MB, ..., YB までの
// const 宣言を書きなさい。
package main

import (
	"fmt"
)


const(
	KB = 1000
	MB = KB*KB
	GB = KB*MB
	TB = KB*GB
	PB = KB*TB
	EB = KB*PB
	ZB = KB*EB
	YB = KB*ZB
)

func main(){
	fmt.Printf("1KB = %vByte\n", KB)
	fmt.Printf("1MB = %vByte\n", MB)
	fmt.Printf("1GB = %vByte\n", GB)
	//fmt.Printf("1TB = %vByte\n", TB)  // 以下はオーバーフロー
	//fmt.Printf("1PB = %vByte\n", PB)
	//fmt.Printf("1EB = %vByte\n", EB)
	//fmt.Printf("1ZB = %vByte\n", ZB)
	//fmt.Printf("1YB = %vByte\n", YB)
}
