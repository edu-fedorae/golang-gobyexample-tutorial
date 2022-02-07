package main

import "fmt"

var g int = 3
//h := 4 //Cannot use shorthand syntax outside of function body.

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt. Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(g)

}
