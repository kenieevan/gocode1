package main

import "fmt"

type bailout struct {
	ret int
}

func g() {

	defer func() {

		defer func() {
			p := recover()
			fmt.Printf("In defer defer recover %v\n", p)
		}()

		fmt.Printf("Hello, in defer")
		panic("Panic in defer")
	}()

	fmt.Printf("Hello, in g()")

}

func f() {

	defer func() {
		p := recover()
		if p != nil {
			fmt.Printf("recover %v \n ", p)
			panic("panic from defer recover")
		}
	}()
	//defer fmt.Println("1")
	//defer fmt.Println("2")

	panic("panic from f")
}

func deferval(ret *int) {
	p := recover()
	*ret = 1
	fmt.Printf("%v \n", p)
	return
}

func testf() (result int) {
	defer deferval(&result)
	panic(bailout{1})
}

func main() {
	ret := testf()
	fmt.Println(ret)

	g()
}
