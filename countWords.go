package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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

	input := bufio.NewScanner(resp.Body)
	input.Split(bufio.ScanWords)
	wordCnt := 0
	imgCnt := 0

	for input.Scan() {
		if strings.Contains(input.Text(), "<img") || strings.Contains(input.Text(), "<IMG") {
			imgCnt++
		}
		fmt.Printf("%s\n", input.Text())
		wordCnt++
	}

	fmt.Printf("-- word count is %d  img cnt is %d ---\n", wordCnt, imgCnt)

	resp.Body.Close()
}
