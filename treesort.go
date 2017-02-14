package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//sort in place. the values is slices
func sort(values []int) {
	//insert the values into a tree start from root
	var root *tree

	for _, v := range values {
		root = insertValue(root, v)
	}

	//then traverse the tree, and sort the values into slices
	appendValue(values[:0], root)

}

func insertValue(root *tree, v int) *tree {

	if root == nil {
		root = &tree{value: v, left: nil, right: nil}
		return root
	}

	if v < root.value {
		root.left = insertValue(root.left, v)
	} else {
		root.right = insertValue(root.right, v)
	}

	return root
}

func appendValue(values []int, t *tree) []int {
	if t != nil {

		values = appendValue(values, t.left)
		values = append(values, t.value)
		values = appendValue(values, t.right)
	}

	return values
}

func main() {
	values := []int{199, 5, 8, 1, 9, 0}
	sort(values)
	fmt.Println(values)
}
