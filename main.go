package main

import (
	_ "errors"
	"fmt"
	_ "unicode/utf8"
)

func main() {
	// var intNum rune = 'a'
	// 	var numerator int = 88
	// 	var denominator int = 0
	// 	var result, remainder, err = intDivision(numerator, denominator)
	// 	// if err != nil {
	// 	// 	fmt.Printf(err.Error())
	// 	// } else if remainder == 0 {
	// 	// 	fmt.Printf("The result of the division is %v", result)
	// 	// } else {
	// 	// 	fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	// 	// }
	//
	// 	//switch statement
	// 	switch {
	// 	case err != nil:
	// 		fmt.Printf(err.Error())
	// 	case remainder == 0:
	// 		fmt.Printf("The Result of the int division is %v", result)
	// 	default:
	// 		fmt.Printf("The result of the int division is %v with remainder: %v", result, remainder)
	// 	}
	// }
	// func intDivision(numerator int, denominator int) (int, int, error) {
	// 	var err error
	// 	if denominator == 0 {
	// 		err = errors.New("Cannot divide by zero")
	// 		return 0, 0, err
	// 	}
	// 	var result int = numerator / denominator
	// 	var remainder int = numerator % denominator
	// 	return result, remainder, err

	//SLICES
	// slice_Name := []string{"Humayoon", "Rehman", "Mujtaba", "Talha"}
	// length, capacity := len(slice_Name), cap(slice_Name)
	// fmt.Printf("The length of the slice is: %v and the capacity of the slice is: %v", length, capacity)

	my_slice := make([]int, 5, 10)
	for i := 0; i < 5; i++ {
		my_slice[i] = (i + 1) * 10
	}
	for i, v := range my_slice {
		fmt.Printf("Index %d: %d\n", i, v)
	}
	for i := range my_slice {
		my_slice = append(my_slice, (i+6)*10)
	}

	fmt.Println("index after appending")

	for i, v := range my_slice {
		fmt.Printf("Index %d: %d\n", i, v)
	}
	// var myMap = map[string]uint8{"adam": 22, "Humayoon": 23}
	// fmt.Println(myMap["Humayoon"])
	// var age, ok = myMap["Imran"]
	// if ok {
	// 	fmt.Printf("The age is %v", age)
	// } else {
	// 	fmt.Printf("Invalid Key")
	// }
	//ITERATE THROUGH MAPS OR SLICE OR EVEN ARRAYS
	// for name := range myMap {
	// 	fmt.Printf("Name: %v\n", name)
	// }
}
