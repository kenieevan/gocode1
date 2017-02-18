package main

import "fmt"

func main() {
	s := make([]int, 2, 4)
	fmt.Printf("len is %d cap is %d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("len is %d cap is %d\n", len(s), cap(s))

	var s1 []int
	s1 = make([]int, 2, 4)

	fmt.Printf("len is %d cap is %d\n", len(s1), cap(s1))

	fmt.Println("vim-go")
}
