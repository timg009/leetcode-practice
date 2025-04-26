package main

import "fmt"

func main() {

	numlist := []int{7, 2, 11, 15}
	target := 9

	res := twoSumBrute(numlist, target)
	fmt.Println(res)
}

// 2SumSortedInput
func twoSumBrute(nums []int, target int) []int {
	// O(n^2)
	for i, inum := range nums {
		//fmt.Printf("i %d. num %d\n", i, num)
		for j, jnum := range nums {
			if (inum + jnum) == target {
				return []int{i, j}
			}
		}
	}
	fmt.Printf("target %d.", target)
	return []int{}
}
