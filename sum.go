package main

import "fmt"

func maxMin(first int, vals ...int) (max, min int) {
	max, min = first, first
	for _, val := range vals {
		if max > val {
			max = val
		}

		if min < val {
			min = val
		}
	}
	return
}

func main() {

	fmt.Println(maxMin(1, 4, 0, 8, -4))
	fmt.Println(maxMin(1))
}
