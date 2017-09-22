// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"github.com/waman/exercise-go/ch5/links"
	"log"
	"os"
)

// breadthFirst は worklist 内の個々の項目に対して f を呼び出します。
// f から返された全ての項目は worklist へ追加されます。
// f は、それぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
