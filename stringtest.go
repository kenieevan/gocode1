package main

import "fmt"

func main() {
	s := `Go is a tool for managing Go source code.
  \t 
   usage:   
      go command [arguments]
  `
	fmt.Println(s)

	s1 := "a\ta"
	fmt.Println(s1)

	//s2 := "\xe4\xb8\x96\xe7\x95\x8c"

	h := 0xe4b896

	fmt.Printf("%x", h)
}
