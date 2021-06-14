package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] 			// creating a slice- a[start:end] - treated ass --> a[start:end-1]
	fmt.Println(b)					// slice doesnt own any data
}