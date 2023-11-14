package main

import (
	"fmt"
)

// Binary Search accepts an Slice of ordered Uints and a element returns the index of that element element.
func bSearch(data []uint, elm uint) int {
	var l, r, m int = 0, len(data) - 1, 0
	for data[m] != elm {
		m = (l + r) / 2
		fmt.Printf("l = %v, r = %v, m = %v, data[m] = %v\n", l, r, m, data[m])
		if elm < data[m] {
			r = m
		} else {
			l = m
		}
	}
	return m
}

func main() {
	var data = []uint{
		4,
		5,
		14,
		24,
		57,
		98,
		102,
		125,
		150,
		198,
		201,
		205,
		208,
		210,
		245,
	}
	var elm uint = 150
	index := bSearch(data, 150)
	fmt.Printf("The index of %v is %v\n", elm, index)
}
