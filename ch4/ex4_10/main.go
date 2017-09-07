// 【練習問題 4.10】
// 一か月未満、一年未満、一年以上の期間で分類された結果を
// 報告するように issues を修正しなさい。
package main

import (
	"os"
	"log"
	"fmt"
	"github.com/waman/exercise-go/ch4/github"
	"time"
)

type Period string

const(
  WithinAMonth Period = "一か月未満"
	WithinAYear         = "一年未満"
	OverAYear           = "一年以上"
)

// 実行例：
//
//   > go run ./ch4/ex4_10/main.go repo:golang/go is:open json decoder
//
func main(){
	var categ = map[Period][]github.Issue{
		WithinAMonth: {},
		WithinAYear : {},
		OverAYear   : {},
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		h := time.Since(item.CreatedAt).Hours()
		switch{
		case h < 24*30:  // 1か月 = 30日
			categ[WithinAMonth] = append(categ[WithinAMonth], *item)

		case h < 24*365:  // 1年 = 365日
			categ[WithinAYear] = append(categ[WithinAYear], *item)

		default:
			categ[OverAYear] = append(categ[OverAYear], *item)
		}
	}

	printIssuesInPeriod(categ, WithinAMonth)
	printIssuesInPeriod(categ, WithinAYear)
	printIssuesInPeriod(categ, OverAYear)
}

func printIssuesInPeriod(categ map[Period][]github.Issue, period Period){
	issues := categ[period]

  fmt.Printf("***** %s *****\n", period)
	fmt.Printf("%d issues:\n", len(issues))

	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
