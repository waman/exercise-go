// 【練習問題 4.11】
// コマンドラインからユーザが GitHub のイシューを作成、読み出し、
// 更新、クローズできるツールを構築しなさい。大量のテキストを入力
// する必要がある場合には、ユーザの好みのテキストエディタを起動する
// ようにしなさい。
package main

import (
	"time"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string  // マークダウン形式
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// GetIssues は GitHub のイシューを取得します
func GetIssues(repo string){
	getURL := fmt.Sprintf("https://api.github.com/repos/%s/issues", repo)
	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("イシューを取得できませんでした %s", resp.Status)
		return
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// CreateIssue は GitHub のイシューを作成します。
func CreateIssue(repo, title, message string){
	createURL := fmt.Sprintf("https://api.github.com/repos/%s/issues?title=%s", repo, title)
	fmt.Println(createURL)

	issue := Issue{Title:title, Body:message}
	_, err := json.Marshal(issue)
	if err != nil {
		log.Fatal(err)
		return
	}

	body := `{
	  "title":"IssueTitle",
	  "body":"Issue Body",
	  "assignee":"waman",
	  "milestone":1,
	  "labels":["test"]
	}`

	resp, err := http.Post(createURL, "application/json", strings.NewReader(body))
	//resp, err := http.Post(createURL, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("イシューを作成できませんでした： %s", resp.Status)
		return
	}

	fmt.Println(resp.Body)
	resp.Body.Close()
}

// 実行例：
//
//   > go run ./ch4/ex4_11/main.go waman/exercise-go create "Issue Title" "Issue Body"
//   > go run ./ch4/ex4_11/main.go waman/exercise-go get
//
func main(){
	switch os.Args[2]{
	case "create":
		CreateIssue(os.Args[1], os.Args[3], os.Args[4])
	case "get":
		GetIssues(os.Args[1])
	}
}