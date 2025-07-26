package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			ans := []int{j, i}
			fmt.Println("answer:", ans) // print here
			fmt.Println("nums/target:", nums, target)
			return ans
		}
		seen[n] = i
	}
	fmt.Println("NO Solution? nums/target:", nums, target)
	return nil
}
