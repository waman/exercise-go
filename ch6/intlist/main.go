// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import "fmt"

// IntList は整数のリンクリストです。
// nil の *IntList は空リストを表します。
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum はリスト要素の合計値を返します。
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	list := &IntList{1, &IntList{2, &IntList{3, nil}}}
	fmt.Println(list.Sum()) // 6
}
