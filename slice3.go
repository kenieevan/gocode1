package main

import "fmt"

func noempty2(strings []string) []string {

	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func main() {
	data := []string{"1", "", "2"}
	out := noempty2(data)
	fmt.Println(out)
	fmt.Println(data)
}
