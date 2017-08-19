package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	counts := make(map[string]int)
	containingFiles := make(map[string]string)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		localCounts := make(map[string]int)
		for input.Scan() {
			s := input.Text()
			counts[s]++
			localCounts[s]++
			if localCounts[s] == 1 {
				containingFiles[s] += "[" + arg + "]"
			}
		}
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s %s\n", n, line, containingFiles[line])
		}
	}
}
