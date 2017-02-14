package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Set(val int) {

	whichByte := val / 64
	whichBit := uint(val % 64)

	for whichByte >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[whichByte] |= 1 << whichBit
}

func (s *IntSet) String() string {

	var buf bytes.Buffer
	for i, word := range s.words {
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				fmt.Fprintf(&buf, "%d ", i*64+bit)
			}
		}
	}
	return buf.String()
}

func main() {

	var x IntSet
	(&x).Set(10)
	fmt.Println(x.String())
	fmt.Println(&x)      //ok, invoke the String
	fmt.Printf("%s", &x) //
	fmt.Printf("%s", x)  // no string

}
