package main

import (
	"fmt"
	//	"sort"
)

var prereqs = map[string][]string{
	"alg":    {"ds"},
	"cal":    {"linear"},
	"com":    {"ds", "formal", "org"},
	"ds":     {"dm"},
	"db":     {"ds"},
	"dm":     {"intro"},
	"formal": {"dm"},
	"net":    {"os"},
	"os":     {"ds", "org"},
	"prog":   {"ds", "org"},
}

func topoSort(m map[string][]string) []string {
	//sort the keys
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	//sort.Strings(keys)

	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)

	visitAll = func(items []string) {

		for _, item := range items {
			if seen[item] != true {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	visitAll(keys)
	return order

}

func main() {

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
