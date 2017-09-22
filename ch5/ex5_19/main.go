// 【練習問題 5.19】
// return 文を含んでいないのに、ゼロ値ではない値を返す関数を
// panic と recover を使って書きなさい。
package main

func main() {
	println(f())
}

func f() (result int) {
	type bailout struct{ x int }

	defer func() {
		switch p := recover(); p {
		case 10:
			result = 10
		}
	}()

	panic(10)
}
