package main

import "fmt"

var testmap = make(map[string]int)
var testmap1 = map[int]int{1: 1, 2: 2}

func main() {
	a := []int{1, 2, 4}
	var a1 []int = []int{1, 2, 4}
	//var a1 []int = {1, 2, 4}
	fmt.Printf("%v %v", a, a1)
	testmap["1"] = 1
	fmt.Printf("%#v \n", testmap1)
	for k, v := range testmap1 {
		fmt.Printf("key : %d value %d \n", k, v)
	}
	fmt.Println("vim-go")
}
