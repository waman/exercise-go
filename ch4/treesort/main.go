package main

import (
	"os"
	"strconv"
	"fmt"
)

// コマンドライン引数から1つ以上の int 値をとって Sort() によってソートする
func main(){
	n := len(os.Args[1:])
	if n == 0 {
		fmt.Println("1つ以上の整数値の引数が必要です。")
		return
	}

	var ints []int
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("数値でない引数が現れました: %s\n\t%v\n", arg, err)
			return
		}
		ints = append(ints, i)
	}

	Sort(ints)
	fmt.Println(ints)
}

type tree struct {
	value       int
	left, right *tree
}

// Sort は values 内の値をその中でソートします。
func Sort(values []int){
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	}else{
		t.right = add(t.right, value)
	}
	return t
}

// appendValues は t の要素を values の正しい順序に追加し、
// 結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
