package main

import (
	"testing"
)

func Test0個の引数を許すmaxで1個以上の引数の場合(t *testing.T) {
	result, err := max(1, 3, 7, 2, 5)
	if err != nil{
		t.Error("最大値を計算できる場合にエラーが返されました。")
	}

	if result != 7 {
		t.Errorf("最大値が間違っています： %d", result)
	}
}

func Test0個の引数を許すmaxで0個の引数の場合(t *testing.T){
	_, err := max()
	if err == nil {
		t.Error("0個の引数を渡したのにエラーが返されませんでした。", err)
	}
}

func Test少なくとも一つ以上の引数が必要なMax(t *testing.T){
	//mx3 := Max()  // コンパイルエラー
	result := Max(1, 3, 7, 2, 5)
	if result != 7 {
		t.Errorf("結果が違います： %d != 7", result)
	}
}

func Test0個の引数を許すminで1個以上の引数の場合(t *testing.T) {
	result, err := min(1, 3, 7, 2, 5)
	if err != nil{
		t.Error("最小値を計算できる場合にエラーが返されました。")
	}

	if result != 1 {
		t.Errorf("最小値が間違っています： %d", result)
	}
}

func Test0個の引数を許すminで0個の引数の場合(t *testing.T){
	_, err := min()
	if err == nil {
		t.Error("0個の引数を渡したのにエラーが返されませんでした。", err)
	}
}

func Test少なくとも一つ以上の引数が必要なMin(t *testing.T){
	//mx3 := Max()  // コンパイルエラー
	result := Min(1, 3, 7, 2, 5)
	if result != 1 {
		t.Errorf("結果が違います： %d != 7", result)
	}
}