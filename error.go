package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func main() {

	var e error = syscall.ENOENT
	_, err := os.Open("/nosuchfile")

	var e1 = errors.New("1")
	var e2 = errors.New("1")

	if e1 == e2 {
		fmt.Println("equal")
	} else {

		fmt.Println("not equal")
	}

	fmt.Printf("%#v %#v \n", err, e)
}
