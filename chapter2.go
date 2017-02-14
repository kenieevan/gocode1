package main

import "fmt"

const (
	freezeF = 0
	boilF   = 212
)

type Celsius float64
type Fahrenheit float64

var floatVar = 0.2

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) String() string { return fmt.Sprintf("%g degree", c) }

//it seems syntax error to invoke function in short var declare
func main() {
	//	var BoilingF Fahrenheit = CToF(100)
	//	var FreezingF Fahrenheit = CToF(0)

	//	fmt.Printf("%g\n", BoilingF-FreezingF)

	fmt.Println(AbsoluteZeroC.String()) //degree
	fmt.Printf("%v\n", AbsoluteZeroC)   //degree
	fmt.Printf("%s\n", AbsoluteZeroC)   //degree
	fmt.Printf("%q\n", AbsoluteZeroC)   //degree
	fmt.Println(AbsoluteZeroC)          //degree

	fmt.Println(float64(AbsoluteZeroC)) //-273.15
	fmt.Printf("%g\n", AbsoluteZeroC)   //-273.15 floating point number

}
