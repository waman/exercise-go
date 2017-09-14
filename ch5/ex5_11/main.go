package main

import (
	"fmt"
	"sort"
	"log"
)

// prereqs は情報科学の各講座をそれぞれの事前条件となる講座と対応付けします。
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra": {"calculus"},  // 追加の依存
}

func main(){
	for i, course := range topoSort(prereqs){
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

type state int
const(
	notAppended  state = iota  // ゼロ値
	nowAppending
	appended
)

func topoSort(m map[string][]string) []string {
	var order []string
	entryStates := make(map[string]state)
	var visitAll func (items []string, entryStates map[string]state)

	visitAll = func(items []string, entryStates map[string]state){
		for _, item := range items {
			switch entryStates[item] {
			case notAppended:
				entryStates[item] = nowAppending
				visitAll(m[item], entryStates)
				entryStates[item] = appended
				order = append(order, item)

			case nowAppending:
				log.Fatalf("循環した依存性が出現しました： %s", item)

			//case appended:
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys, entryStates)
	return order
}
