// 【練習問題 4.3】
// スライスの代わりに配列へのポインタを使うように reverse を書き直しなさい。
package main

import "fmt"

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
