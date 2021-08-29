package main

import "fmt"

type ArrProcessingFunc func([]int) []int

func main() {
	fmt.Println("--- task04 ---")

	fmt.Println("\n-- first try --")
	OutputResults([]int{17, 18, 5, 4, 6, 1}, NextGreatestReplace)
	OutputResults([]int{400}, NextGreatestReplace)

	fmt.Println("\n-- second try --")
	OutputResults([]int{17, 18, 5, 4, 6, 1}, NextGreatestReplaceV2)
	OutputResults([]int{400}, NextGreatestReplaceV2)
}

func OutputResults(arr []int, procfunc ArrProcessingFunc) {
	fmt.Println("\nInput: arr =", arr)
	fmt.Println("Output:", procfunc(arr))
}

// My first try. Left-to-Right. Made in 8 minutes. Complexity O(n^2)
func NextGreatestReplace(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var max int
		for _, el := range arr[i+1:] {
			if max < el {
				max = el
			}
		}
		arr[i] = max
	}
	arr[len(arr)-1] = -1
	return arr
}

// My second try. Right-to-Left. Complexity O(n)
func NextGreatestReplaceV2(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = max
		if max < temp {
			max = temp
		}
	}
	return arr
}
