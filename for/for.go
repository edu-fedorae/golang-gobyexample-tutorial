package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println("I:",i)
		i++
	}

	for j := 8; j <= 9; j++ {
		fmt.Println("J:", j)
	}

	for {
		fmt.Println("Looping...")
		break // or return
	}

	for n := 0; n <= 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
