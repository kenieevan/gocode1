package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {

	wordCnt := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		wordCnt++
	}
	*c = WordCounter(wordCnt)
	return wordCnt, nil
}

func main() {

	var c WordCounter

	fmt.Fprintf(&c, "hello world 111 222")
	fmt.Println(c)
}
