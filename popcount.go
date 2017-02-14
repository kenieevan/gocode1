package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	result := 0
	var i uint32
	for i = 0; i < 8; i++ {
		// >> * have the same priority... so the () is a must here!!!
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

func main() {

	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("string conv failed")
	}
	var v1 uint64 = uint64(v)
	fmt.Println(PopCount(v1))
}
