package main

import "fmt"

func main() {

	// here increment is done in loop body
	var i int = 1
	for i <= 4 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		// here increment is done in the for loop
		fmt.Println("#", j)
	}

	// Range in for loop ( it starts with zero iteration)

	for x := range 3 {
		fmt.Println("Deovrath", x)
	}

	// using break and continue in for loop

	for s := range 20 {
		if s%2 == 0 {
			if s == 12 {
				fmt.Println("Break at", s)
				break
			} else {
				fmt.Println("even no", s)
				continue
			}

		}
	}
}
