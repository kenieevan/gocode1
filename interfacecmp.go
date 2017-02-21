package main

import "fmt"

type testcmp interface {
	t1() int
}

type s struct {
	i int
}

func (s1 s) t1() int {
	return s1.i
}

type t struct {
	i int
}

func (t1 t) t1() int {
	return t1.i
}

func main() {

	//the dynamic type t and s is not equal...
	var ti1 testcmp = s{1}
	var ti2 testcmp = s{2}

	if ti1 == ti2 {
		println("interface cmp ok")
	} else {
		println("interface cmp not equal")
	}

	ti1, ok := ti2.(testcmp) //the value change after assignment
	fmt.Printf("assin ok? %b   value is %#v\n", ok, ti1)
}
