// 【練習問題 4.11】
// コマンドラインからユーザが GitHub のイシューを作成、読み出し、
// 更新、クローズできるツールを構築しなさい。大量のテキストを入力
// する必要がある場合には、ユーザの好みのテキストエディタを起動する
// ようにしなさい。
package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"strings"
	"io/ioutil"
	"io"
	"github.com/waman/exercise-go/ch4/github"
	"net/url"
	"strconv"
)

var accessToken = getAccessToken()

func getAccessToken() string {
	bs, err := ioutil.ReadFile(".github_access_token")
	if err != nil {
		log.Fatal("アクセストークンの読み込みに失敗しました： .github_access_token")
		os.Exit(1)
	}
	return string(bs)
}

func createRequest(method, urlStr string, body io.Reader) (*http.Response, error){
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil { return nil, err }

	req.Header.Add("Authorization", "token " + accessToken)

	client := &http.Client{}
	return client.Do(req)
}

// create, update, close が成功したときに、メッセージと対象 URL を表示します。
func printResult(resp *http.Response, successMessage string){
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var issueURL struct{URL string}
	if err := json.Unmarshal(body, &issueURL); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s： %s\n", successMessage, issueURL.URL)
}

// 自分の GitHub アカウント情報を取得します。　デフォルトの動作。
func user(){
	resp, err := createRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Fatalf("アカウント情報を取得できませんでした： %s", err)
		return
	}
	defer resp.Body.Close()  // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("アカウント情報を取得できませんでした： %s", resp.Status)
		return
	}

	var user github.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Login: %s\nURL  : %s\n", user.Login, user.HTMLURL)
}

// CreateIssue は GitHub のイシューを作成します。
func CreateIssue(repo, title, message string){
	createIssuesURL := fmt.Sprintf("https://api.github.com/repos/%s/issues?title=%s",
		                             repo, url.QueryEscape(title))
	issueBody := fmt.Sprintf("{\"title\":\"%s\", \"body\":\"%s\"}", title, message)

	resp, err := createRequest("POST", createIssuesURL, strings.NewReader(issueBody))
	if err != nil {
		log.Fatalf("イシューを作成できませんでした： %s", err)
		return
	}
	defer resp.Body.Close()  // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("イシューを作成できませんでした： %s", resp.Status)
		return
	}

	printResult(resp, "イシューを作成しました")
}

// GetIssues は GitHub のイシューを取得します
func GetIssues(repo string){
	getIssuesURL := fmt.Sprintf("https://api.github.com/repos/%s/issues", repo)
	resp, err := createRequest("GET", getIssuesURL, nil)
	if err != nil {
		log.Fatalf("イシューを取得できませんでした： %s", err)
		return
	}
	defer resp.Body.Close()  // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("イシューを取得できませんでした： %s", resp.Status)
		return
	}

	var issues []github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// UpdateIssue は GitHub のイシューを更新します。
func UpdateIssue(repo, no, title, message string){
	if _, err := strconv.Atoi(no); err != nil {
		log.Fatal(err)
		return
	}

	updateIssuesURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%s", repo, no)
	issueBody := fmt.Sprintf("{\"title\":\"%s\", \"body\":\"%s\"}", title, message)

	resp, err := createRequest("PATCH", updateIssuesURL, strings.NewReader(issueBody))
	if err != nil {
		log.Fatalf("イシューを更新できませんでした： %s", err)
		return
	}
	defer resp.Body.Close()  // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("イシューを更新できませんでした： %s", resp.Status)
		return
	}

	printResult(resp, "イシューを更新しました")
}

// CloseIssue は GitHub のイシューをクローズします。
func CloseIssue(repo, i string){
	if _, err := strconv.Atoi(i); err != nil {
		log.Fatal(err)
		return
	}

	closeIssuesURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%s", repo, i)
	issueBody := "{\"state\":\"closed\"}"

	resp, err := createRequest("PATCH", closeIssuesURL, strings.NewReader(issueBody))
	if err != nil {
		log.Fatalf("イシューをクローズできませんでした： %s", err)
		return
	}
	defer resp.Body.Close()  // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("イシューをクローズできませんでした： %s", resp.Status)
		return
	}

	printResult(resp, "イシューをクローズしました")
}

// GitHub のアクセストークンを作業ディレクトリ上のファイル ".github_access_token" に保存して
// おいてください。　アクセストークンの生成は https://github.com/settings/token からできます。
// 入力が長い場合のエディタの起動は実装していません。
//
// 自分のアカウントの確認：
//
//   > go run ./ch4/ex4_11/main.go
//
// イシューの作成：
//
//   > go run ./ch4/ex4_11/main.go create waman/exercise-go "Issue Title" "Issue Body"
//
// イシューの取得：
//
//   > go run ./ch4/ex4_11/main.go get waman/exercise-go
//
// イシューの更新：
//
//   > go run ./ch4/ex4_11/main.go update waman/exercise-go 1 "New Issue Title" "New Issue Body"
//
// イシューのクローズ：
//
//   > go run ./ch4/ex4_11/main.go close waman/exercise-go 1
//
func main(){

	if len(os.Args) < 2 {
		user()
		return
	}

	// os.Args[2] -> repo
	ss := strings.Split(os.Args[2], "/")
	repo := url.QueryEscape(ss[0]) + "/" + url.QueryEscape(ss[1])

	switch os.Args[1]{
	case "create":
		CreateIssue(repo, os.Args[3], os.Args[4])
	case "get":
		GetIssues(repo)
	case "update":
		UpdateIssue(repo, os.Args[3], os.Args[4], os.Args[5])
	case "close":
		CloseIssue(repo, os.Args[3])
	default:
		user()
	}
}