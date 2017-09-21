// 【練習問題 5.14】
// 異なる構造を調べるために breadthFirst 関数を使いなさい。たとえば、topoSort
// の例（有向グラフ）の講座の依存関係、コンピュータ上のファイルシステムの階層、
// 公共機関のウェブサイトからダウンロードしたバスや地下鉄の経路（無向グラフ）の
// リストなどを利用できます。
package main

import (
	"fmt"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"os"
)

// 無向グラフは循環を除く処理を行わないといけないで教育的だと思いますが、
// ここでは扱いません。
func main() {
	fmt.Println("講座の依存関係の解析")
	analyzeClasses()
	fmt.Println()

	fmt.Println("ディレクトリ階層の解析（カレントディレクトリ）")
	analyzeDirectoryStructure()
}

// topoSort で使った講座の依存関係
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func analyzeClasses() {
	var keys = make([]string, len(prereqs))
	for key := range prereqs {
		keys = append(keys, key)
	}

	htmlutil.BreadthFirst(crawlClasses, keys)
}

func crawlClasses(class string) []string {
	fmt.Println(class)
	return prereqs[class]
}

func analyzeDirectoryStructure() {
	htmlutil.BreadthFirst(crawlFiles, []string{"."})
}

func crawlFiles(filename string) []string {
	// .git 内はファイルがたくさんあるので無視
	if filename == "./.git" {
		return nil
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	defer func() {
		if cErr := file.Close(); err == nil && cErr != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	fi, err := os.Lstat(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	if fi.IsDir() {
		fmt.Printf("[DIR]  %s\n", filename)

		dirs, err := file.Readdirnames(-1)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}

		for i, dir := range dirs {
			dirs[i] = filename + "/" + dir
		}

		return dirs

	} else {
		fmt.Printf("[File] %s\n", filename)
		return nil
	}
}
