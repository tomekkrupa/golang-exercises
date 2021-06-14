package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	s := a[1:4]
	/*
	s1 := a[1:4]
	fmt.Println("original array", a)
	fmt.Println("sliced array", s)
	fmt.Println("sliced array s1", s1)

	for i := range s {
		s[i]++						// increment every number in array by 1
	}

	fmt.Println("original array", a)
	fmt.Println("sliced array", s)
	fmt.Println("sliced array s1", s1)	

	for i := range s1 {
		s1[i]++						// increment every number in array by 1
	}

	fmt.Println("original array", a)
	fmt.Println("sliced array", s)
	fmt.Println("sliced array s1", s1)
	
	fmt.Println(len(s))
	fmt.Println(cap(s))
	*/

	s = s[:cap(s)] 					// re-slicing s up to its capacity --> note the last println which shows whats inside of the array
	fmt.Println(a)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)
}

// analysis of the output: after incrementing elements through slice s and s1 ---> NOTE: original array at each iteration will 
// also increase by 1
// NOTE: each slice modifies the original array

// original array [1 2 3 4 5 6 7]
// sliced array [2 3 4]
// sliced array s1 [2 3 4]
// original array [1 3 4 5 5 6 7]
// sliced array [3 4 5]
// sliced array s1 [3 4 5]
// original array [1 4 5 6 5 6 7]
// sliced array [4 5 6]
// sliced array s1 [4 5 6]