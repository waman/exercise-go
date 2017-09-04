// パッケージ github は、GitHub のイシュートラッカーに対する
// Go の API を提供します。
//
//   http://developer.github.com/v3/serch/#search-issues
//
// を参照のこと。
package github

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
)

const IssuesURL = "https://api.github.com/serch/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

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

// SearchIssues は GitHub のイシュートラッカーに問い合わせます。
func SearchIssues(terms []string) (*IssueSearchResult, error){
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}