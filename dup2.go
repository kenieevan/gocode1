package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range filenames {
			fd, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(fd, counts)
			fd.Close()

		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s \n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

}
