package main

import "fmt"

func stacktest() *int {
	testval := 100
	return &testval
}

func stacktest1() (*int, *string) {
	testval := 101
	testval1 := 101
	testval2 := 101
	testval3 := 101
	testval4 := 102
	testval5 := 102
	testval6 := 103
	testval7 := 104

	teststr := "hello"
	teststr2 := "hello"

	fmt.Println(testval, testval1, testval2, testval3, testval4, testval5, testval6, testval7, teststr2)
	return &testval, &teststr
}

func main() {
	pInt := stacktest()
	fmt.Printf("1: %d\n", *pInt)
	stacktest1()
	stacktest1()
	stacktest1()
	fmt.Printf("2: %d\n", *pInt)

}
