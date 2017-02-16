package main

import "fmt"

type Track struct {
	Title string
}

var tracks = []*Track{{"GO"}}
var t = []*struct{ int }{{1}, {2}}

func main() {
	fmt.Printf("%q\n", tracks)
	fmt.Printf("%q\n", t)
}
