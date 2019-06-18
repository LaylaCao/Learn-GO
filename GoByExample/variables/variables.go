package main

import "fmt"

func main()  {
	var a = "initial"
	fmt.Println(a)

	b,c  := 1,2
	//var b,c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)


	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)
}