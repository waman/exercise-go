// 【練習問題 3.12】
// 二つの文字列が互いにアナグラムになっているか、すなわち
// 同じ文字を異なる順番で含んでいるかを報告する関数を書きなさい。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// コマンドライン引数（なければ標準入力）で与えられた文字列に isAnagram を
// 適用して結果を表示する
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)

		fmt.Println("スペースで区切って2つの文字列を入力してください")
		fmt.Print("> ")
		for input.Scan() {
			ss := strings.Split(input.Text(), " ")
			fmt.Println(isAnagram(ss[0], ss[1]))

			fmt.Println("スペースで区切って2つの文字列を入力してください")
			fmt.Print("> ")
		}
	} else {
		fmt.Println(isAnagram(args[1], args[2]))
	}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, r := range s1 {
		if i := strings.IndexRune(s2, r); i >= 0 {
			s2 = s2[:i] + s2[i+1:]
		} else {
			return false
		}
	}

	return len(s2) == 0
}
