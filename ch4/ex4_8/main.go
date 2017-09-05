// 【練習問題 4.8】
// unicode.IsLetter などの関数を使って、Unicode 分類に従って
// 文字や数字などを数えるように charcount を修正しなさい。
package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"unicode"
)

type Category string

const(
	Letter  Category = "L"
	Number           = "N"
	Symbol           = "(Symbol)"
	Punct            = "P"
	Mark             = "M"
	Space            = "(Space)" // unicode.IsSpace() で判定するのでカテゴリZは異なる
	Unknown          = "?"
)

// 実行例
//
//   > go run ./ch4/ex4_8/main.go < ./ch4/TheGoBlog-strings.txt
//
// go get でコードを取得した場合は、上記のコマンドではうまく動かないかもしれません。
func main(){
	counts := make(map[Category]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r){
			counts[Letter]++
		}else if unicode.IsNumber(r){
			counts[Number]++
		}else if unicode.IsSymbol(r){
			counts[Symbol]++
		}else if unicode.IsMark(r){
			counts[Mark]++
		}else if unicode.IsPunct(r){
			counts[Punct]++
		}else if unicode.IsSpace(r){
			counts[Space]++
		}else{
			fmt.Printf("%q ", r)
			counts[Unknown]++
		}
	}
	fmt.Printf("\ncategory\tcount\n")

	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
