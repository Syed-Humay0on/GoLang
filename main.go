package main

import (
	_ "errors"
	"fmt"
	_ "unicode/utf8"
)

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// go twoSum(nums, target)
	// time.Sleep(time.Second * 2)
	// fmt.Println("Hello Web")
	names := GetNames()
	fmt.Println(names)
}
