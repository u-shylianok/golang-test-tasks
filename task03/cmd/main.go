package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- task03 ---")

	// task02
	rockets := CreateSomeRockets()
	for i, rocket := range rockets {
		fmt.Println(i, "|", rocket)
	}

	var movers [len(rockets)]Mover
	for i, rocket := range rockets { // this anti optimization is ONLY for code separation
		movers[i] = rocket
	}

	// task01 with array
	for _, mover := range movers {
		mover.Move()
	}

}
