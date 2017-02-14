package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//just copy the strings.Reader.go
type MyReader struct {
	s string
	i int
}

//implement the io.Reader interface
// b is the output buffer
func (r *MyReader) Read(b []byte) (n int, err error) {
	//copy value in string to b, the caller shall allocate b...
	// return n bytes copied
	if len(b) == 0 {
		return 0, nil
	}

	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:]) // if len of b is not enough, the copy will hande it...
	r.i += n
	return
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		} //for
	} //if

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
		os.Exit(1)
	}

	var ReadPtr *MyReader = &MyReader{s: string(b)}
	doc, err := html.Parse(ReadPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink %v \n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}
