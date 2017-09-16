// net/url パッケージの Values 型の定義の一部
package main

import "fmt"

// Values は文字列キーを値のリストに対応付けします。
type Values map[string][]string

// Get は指定されたキーに関連づけられた最初の値を返します。
// あるいは、値がなければ "" を返します。
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add は値をキーに追加します。
// キーに関連づけられた既存の値に追加します。
func (v Values) Add(key, value string){
	v[key] = append(v[key], value)
}

func main(){
	m := Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))  // "en"
	fmt.Println(m.Get("q"))     // ""
	fmt.Println(m.Get("item"))  // "1"
	fmt.Println(m["item"])      // "[1 2]"

	m = nil
	fmt.Println(m.Get("item"))  // ""
	m.Add("item", "3")          // panic!
}
