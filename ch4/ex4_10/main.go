// 【練習問題 4.10】
// 一か月未満、一年未満、一年以上の期間で分類された結果を
// 報告するように issues を修正しなさい。
package main

import (
	"fmt"
	"github.com/waman/exercise-go/ch4/github"
	"log"
	"os"
	"time"
)

type Period string
type Catalog map[Period][]*github.Issue

const (
	WithinAMonth Period = "一か月未満"
	WithinAYear         = "一年未満"
	OverAYear           = "一年以上"
)

// 実行例：
//
//   > go run ./ch4/ex4_10/main.go repo:golang/go is:open json decoder
//
func main() {
	var catalog = Catalog{
		WithinAMonth: {},
		WithinAYear:  {},
		OverAYear:    {},
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		h := time.Since(item.CreatedAt).Hours()
		switch {
		case h < 24*30: // 1か月 = 30日
			catalog[WithinAMonth] = append(catalog[WithinAMonth], item)

		case h < 24*365: // 1年 = 365日
			catalog[WithinAYear] = append(catalog[WithinAYear], item)

		default:
			catalog[OverAYear] = append(catalog[OverAYear], item)
		}
	}

	printIssuesInPeriod(catalog, WithinAMonth)
	printIssuesInPeriod(catalog, WithinAYear)
	printIssuesInPeriod(catalog, OverAYear)
}

func printIssuesInPeriod(catalog Catalog, period Period) {
	issues := catalog[period]

	fmt.Printf("***** %s *****\n", period)
	fmt.Printf("%d issues:\n", len(issues))

	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
