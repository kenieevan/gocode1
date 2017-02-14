package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := []int{0, 1, 2, 3, 4, 5}

	a := [...]int{0, 1, 2, 3, 4, 5}
	a1 := [...]int{0, 1, 2, 3, 4, 5}

	if a == a1 {
		fmt.Println("array equal")
	}
	if s == s1 {
		fmt.Println("slice equal")
	}
	fmt.Println(s)

}
