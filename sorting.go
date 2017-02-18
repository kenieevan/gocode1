package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Length time.Duration
}

var tracks = []*Track{
	{"GO", "Delilah", length("3m38s")},
	{"GO", "Moby", length("3m37s")},
	{"Ready 2 Go", "Martin Solveig", length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

type customerSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customerSort) Len() int           { return len(x.t) }
func (x customerSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customerSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	sort.Sort(customerSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})

	for _, t := range tracks {
		fmt.Printf("%v \n", *t)
	}

}
