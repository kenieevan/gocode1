package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1> {{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
   <th>#</th>
   <th>State</th>
   <th>User</th>
   <th>Title </th>
</tr>
{{range .Items}}
<tr>
   <td><a href='{{.HTMLURL}}' >{{.Number}} </td>
   <td>{{.State}}</td>
   <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
   <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

const IssueURL = "https://api.github.com/search/issues"

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
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {

	q := url.QueryEscape(strings.Join(terms, " "))
	//fmt.Println(IssueURL + "?q=" + q)
	resp, err := http.Get(IssueURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed : %s", resp.Status)
	}

	var result IssueSearchResult
	//Decode(interface).. also used in docker client source code
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}

var result *IssueSearchResult
var err error

func handler(w http.ResponseWriter, r *http.Request) {

	//	fmt.Fprintf(w, "%d issues :\n", result.TotalCount)

	if err := issueList.Execute(w, result); err != nil {
		fmt.Printf("issueList failed %v \n", err)
	}
}

func main() {
	result, err = SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
