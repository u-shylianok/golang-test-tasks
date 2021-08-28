package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- task02 ---")

	rockets := CreateSomeRockets()
	for i, rocket := range rockets {
		fmt.Println(i, "|", rocket)
	}
}
