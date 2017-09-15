// 【練習問題 5.15】
// sum のように、可変個引数関数である max と min を書きなさい。
// 引数なしで呼び出されたら、これらの関数は何をすべきでしょうか。
// 少なくとも一つの引数が必要な別のバージョンも書きなさい。
package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

// 実装のテストは同ディレクトリの main_test.go で行っています。
// テストの実行例は
//
//   > go test ./ch5/ex5_15
//
// 下記の main 関数はコマンドライン引数で指定された整数値から max, min 関数によって
// 最大値、最小値を求めます。　実行例は
//
//   > go run ./ch5/ex5_15/main.go 1 3 7 2 5
//
func main(){
	// コマンドライン引数を int スライスに変換
	var argInts []int
	for _, arg := range os.Args[1:] {
		argInt, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("引数には整数値を指定してください： %s", arg)
		}
		argInts = append(argInts, argInt)
	}

	// 最大値の計算
	mx, err := max(argInts...)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Printf("最大値は %d\n", mx)
	}

	// 最小値の計算
	mn, err := min(argInts...)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Printf("最小値は %d\n", mn)
	}
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("引数は1つ以上指定してください。")
	}

	result := vals[0]
	for _, val := range vals[1:] {
		if val > result {
			result = val
		}
	}
	return result, nil
}

// 少なくとも一つの引数が必要なバージョン
func Max(val0 int, vals ...int) int {
	result := val0
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("引数は1つ以上指定してください。")
	}

	result := vals[0]
	for _, val := range vals[1:] {
		if val < result {
			result = val
		}
	}
	return result, nil
}

// 少なくとも一つの引数が必要なバージョン
func Min(val0 int, vals ...int) int {
	result := val0
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
