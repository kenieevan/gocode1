package main

import "fmt"

//just copy the strings.Reader.go
type Reader struct {
}

//implement the io.Reader interface
// b is the output buffer
func (r *Reader) Read(b []byte) (n int, err error)

func string2Reader() {

}

func main() {
	fmt.Println("vim-go")
}
