package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var sha384 = flag.Bool("s384", false, "sha384")
var s512 = flag.Bool("s512", false, "sha512")

func main() {
	flag.Parse()
	var input string
	var buf bytes.Buffer
	n, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Printf("Get input error %v n is %d\n", err, n)
	}
	fmt.Printf("Input is %s\n", input)
	buf.WriteString(input)
	var b []byte = buf.Bytes()
	c1 := sha256.Sum256(b)
	c3 := sha256.Sum256([]byte(input))
	c2 := sha256.Sum256([]byte("hello,world"))

	fmt.Printf("%x \n %x \n %x\n", c1, c2, c3)
	if *sha384 == true {
		c4 := sha512.Sum384([]byte(input))
		fmt.Printf("sum384 is %x", c4)
	}

	if *s512 == true {
		c5 := sha512.Sum512([]byte(input))
		fmt.Printf("sum512 is %x", c5)
	}

}
