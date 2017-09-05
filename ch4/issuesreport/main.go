package main

import (
	"time"
	"html/template"
	"os"
	"log"
	"github.com/waman/exercise-go/ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
  Funcs(template.FuncMap{"daysAgo": daysAgo}).
  Parse(templ))

// 以下のように実行します：
//
//   > go build ./ch4/issuesreport
//   > issuesreport repo:golang/go is:open json decoder
//
func main(){
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}