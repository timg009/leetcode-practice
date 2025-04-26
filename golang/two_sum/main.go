package main

import "fmt"

func main() {

	numlist := []int{7, 2, 11, 15}
	target := 9

	res := twoSum(numlist, target)
	fmt.Println(res)
}

// 2SumSortedInput
func twoSum(nums []int, target int) []int {

	hm := make(map[int]int)

	for i, num := range nums {
		_, ok := hm[num]
		fmt.Println("ok:", ok)
		if ok {
			return []int{i, hm[num]}
		}

		// Compliment to the target
		hm[target-num] = i
	}

	return []int{}
}

// 2SumSortedInput
func twoSumBrute(nums []int, target int) []int {
	// O(n^2)
	for i, inum := range nums {
		//fmt.Printf("i %d. num %d\n", i, num)
		for j, jnum := range nums {
			if (inum+jnum) == target && i != j {
				return []int{i, j}
			}
		}
	}
	fmt.Printf("target %d.", target)
	return []int{}
}
