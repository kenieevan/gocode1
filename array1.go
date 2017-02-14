package main

import "fmt"

//test slice is acutally passed by copy
func appendInt(x []int, y int) []int {
	zlen := len(x) + 1
	x = x[:zlen]
	x[zlen-1] = y
	return x
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	var x []int = s[0:1]
	var y []int = s[1:2]
	x = y
	fmt.Printf("%v \n", s)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s)
	//x = appendInt(x, 9)
	//fmt.Println(s)
	//fmt.Printf("%d\n", len(x))

}
