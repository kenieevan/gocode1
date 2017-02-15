package main

import (
	"flag"
	"fmt"
)

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

//it shall implent String() and Set() function to satisfy flag.Value() interface
type celsiusFlag struct{ Celsius }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func (c Celsius) String() string { return fmt.Sprintf("%g C", c) }
func F2C(val Fahrenheit) Celsius {
	return Celsius((val - 32) * 5 / 9)
}

//string => 100C  or 200F
func (f *celsiusFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = F2C(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid arg %q", s)
}

func setupCelsiusFlag(name string, defaultVal Celsius, usage string) *Celsius {

	f := celsiusFlag{defaultVal}

	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {

	var temp = setupCelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}

//it seems syntax error to invoke function in short var declare
//func main() {
//	//	var BoilingF Fahrenheit = CToF(100)
//	//	var FreezingF Fahrenheit = CToF(0)
//
//	//	fmt.Printf("%g\n", BoilingF-FreezingF)
//
//	fmt.Println(AbsoluteZeroC.String()) //degree
//	fmt.Printf("%v\n", AbsoluteZeroC)   //degree
//	fmt.Printf("%s\n", AbsoluteZeroC)   //degree
//	fmt.Printf("%q\n", AbsoluteZeroC)   //degree
//	fmt.Println(AbsoluteZeroC)          //degree
//
//	fmt.Println(float64(AbsoluteZeroC)) //-273.15
//	fmt.Printf("%g\n", AbsoluteZeroC)   //-273.15 floating point number
//
//}
