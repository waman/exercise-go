// squares は呼び出されるごとに次の平方数を返す関数を返します。
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main(){
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}