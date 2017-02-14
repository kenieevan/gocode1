package main

import "fmt"

func rmAdjDup(strings []string) []string {
	out := strings[:0]
	var old string

	for _, v := range strings {
		if v != old {
			out = append(out, v)
		}
		old = v
	}
	return out
}

func main() {

	//ages := make(map[string]int)
	//ages := map[string]int{}
	//var ages map[string]int
	var ages map[string]int = map[string]int{}

	ages["ken"] = 31
	for name, age := range ages {
		fmt.Printf("%s %d\n", name, age)
	}
	out := []string{"0", "1", "1", "1", "2", "3", "4", "4"}
	out = rmAdjDup(out)
	fmt.Println(out)
}
