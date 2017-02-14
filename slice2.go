package main

import "fmt"

var runes []rune

func main() {
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q %d\n", runes, len(runes))
}
