// nonempty はスライス内アルゴリズムの例です。
package main

import "fmt"

func main(){
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))  // `["one" "three"]`
	fmt.Printf("%q\n", data)  // `["one" "three", "three"]`

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2))  // `["one" "three"]`
	fmt.Printf("%q\n", data2)  // `["one" "three", "three"]`
}

// nonempty は空文字列ではない文字列を保持するスライスを返します。
// 基底配列は呼び出し中に修正されます。
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}