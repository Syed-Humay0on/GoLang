package main

import (
	_ "errors"
	"fmt"
	_ "unicode/utf8"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans) // [0 1]
}
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
