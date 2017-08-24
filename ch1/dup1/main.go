// dup1 は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
