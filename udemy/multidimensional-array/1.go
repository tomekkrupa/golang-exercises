package main
import "fmt"

func main () {
	a := [3][2]string {{"abc", "def"}, {"xxx", "yyy"}, {"qqq", "zzz"}} //3xrows, each row has got 2x colums
	/*
	fmt.Println(a)
	var b [2][2]string
	b[0][1]="111"
	b[0][0]="000"
	b[1][0]="222"
	b[1][1]="111"
	fmt.Println(b)
	*/
	//use for loop to parse through array
	for _, x := range a { 										//a is a 2xdimensional array- in 1st for loop a gives the row, which is present in x
		fmt.Printf("In 1st for loop = %s\t", x)
		for _, y := range x {									//second for loop will parse through column element present in the row
			fmt.Printf("In 2nd for loop = %s\t", y)
		}
		fmt.Printf("\n") 										// each row separate by new line
	}
}