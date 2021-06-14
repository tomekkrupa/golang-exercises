package main

import "fmt"

func main() {
	a := [2]string{"hello", "world!"}
	fmt.Println(a)
	// [hello world!]
	fmt.Printf("%s\n", a)
	// [hello world!]
	fmt.Printf("%q\n", a)
	// ["hello" "world!"]
}