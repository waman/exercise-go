// 【練習問題 7.3】
// gopl.io/ch4/treesort （4.4節）の *tree 型に対して、
// ツリー内の値の列を見せる String メソッドを書きなさい。
package main

import (
	"os"
	"fmt"
	"strconv"
	"bytes"
)

// コマンドライン引数から1つ以上の int 値をとって Sort() によってソートする
func main() {
	// 引数の処理
	n := len(os.Args[1:])
	if n == 0 {
		fmt.Println("1つ以上の整数値の引数が必要です。")
		return
	}

	var values []int
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("数値でない引数が現れました: %s\n\t%v\n", arg, err)
			return
		}
		values = append(values, i)
	}

	// tree の構築
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	fmt.Println(root.String())
}

func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	t.appendString(&buf)
	buf.WriteString("]")
	return buf.String()
}

func (t *tree) appendString(buf *bytes.Buffer){
	if t.left != nil {
		t.left.appendString(buf)
		buf.WriteString(" ")
	}
	buf.WriteString(strconv.Itoa(t.value))
	if t.right != nil {
		buf.WriteString(" ")
		t.right.appendString(buf)
	}
}

//***** 以下は4.4節より *****
type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
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
