// issues は検索語に一致した GitHub イシューの表を表示します。
package main

import (
	"os"
	"log"
	"fmt"
	"github.com/waman/exercise-go/ch4/github"
)

// 実行例：
//
//   > go run ./ch4/issues/main.go repo:golang/go is:open json decoder
//
func main(){
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
