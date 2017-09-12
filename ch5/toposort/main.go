package main

import (
	"fmt"
	"sort"
)

// prereqs は情報科学の各講座をそれぞれの事前条件となる講座と対応付けします。
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compileers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"descrete math"},
	"databases": {"data sutructures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main(){
	for i, course := range topSort(prereqs){
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func (items []string)

	visitAll = func(items []string){
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
