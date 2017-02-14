package main

import "fmt"

func reverse(ptr *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func rotate(slice []int, shift int) {
	var tmp []int
	l := len(slice)
	for i := 0; i < shift; i++ {
		tmp = append(tmp, slice[i])
	}
	copy(slice[0:], slice[shift:])
	copy(slice[(l-shift):], tmp[:])
}

func main() {
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//reverse(&num)
	rotate(num[:], 2)
	fmt.Println(num)
}
