package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") != true {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
			os.Exit(1)
		}

		for k, v := range resp.Header {
			fmt.Fprintf(os.Stdout, "Header[%q] = %q\n", k, v)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
			os.Exit(1)
		}

		resp.Body.Close()
		fmt.Fprintf(os.Stdout, " %v bytes copied  http status is %v \n", n, resp.Status)
	}
}
