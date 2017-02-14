package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func closefile(f *os.File) (result error) {

	defer func() { result = f.Close() }()
	return nil
}

func fetch(url string) (filename string, n int64, err error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	local := "testdefer"

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	if closeErr := closefile(f); err == nil {
		err = closeErr
	}

	return local, n, err
}

func main() {

	fmt.Println(fetch("http://www.baidu.com"))

}
