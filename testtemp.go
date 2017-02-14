package main

import (
	"fmt"
	"tempconv"
)

//it seems syntax error to invoke function in short var declare
func main() {
	//	var BoilingF Fahrenheit = CToF(100)
	//	var FreezingF Fahrenheit = CToF(0)

	//	fmt.Printf("%g\n", BoilingF-FreezingF)

	fmt.Println(tempconv.AbsoluteZeroC.String()) //degree

}
