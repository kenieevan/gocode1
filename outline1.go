package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func outline(stack []string, n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//XXX depth = depth +1 will get a wrong result, since it will affet the end function
		outline(stack, c, depth+1)
	}

	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: outline url")
		os.Exit(1)
	}
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get error %v \n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "http get url %s error status  is %d\n", url, resp.Status)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http parse error %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()
	outline(nil, doc, 0)
}
