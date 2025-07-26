package main

import (
	_ "errors"
	"fmt"
	_ "unicode/utf8"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
