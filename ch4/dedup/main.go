package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	seen := make(map[string]bool)  // 文字列のセット
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for input.Scan(){
		line := input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Println(line)
			fmt.Print("> ")
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}