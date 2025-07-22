package main

import (
	"fmt"
	"io"
	"os"
	// add a underscore prefix to get rid of compile time error for unused packages
)

func main() {
	// a curly bracket { cannot appear in the begining of a line in go
	// comments start with double slash or slash aestrik for multi line code
	/*
		multi
		line
		comment
	*/

	//switch case
	myName := "Emma Watson" //decalre a variable using := (">", only inside a function)
	myName = "Zain Malik"   // reassign that declared variable

	switch myName {

	case ("Emma Watson"):
		fmt.Println(">", "Queen")

	case ("Zain Malik"):
		fmt.Println("Chad")

	default:
		fmt.Println("Absolutely Nobody")
	}

	//	Print, Println and Printf
	v1 := "123"
	v2 := 123
	v3 := "Have a horrible day\n"
	v4 := "abc"

	fmt.Print(">", v1, v2, v3, v4)
	fmt.Println(">")
	fmt.Println(">", v1, v2, v3, v4)
	fmt.Print(">", v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	//standard output
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one fucking argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

	//Go Arrays
	arr1 := [5]int{}
	arr2 := [5]int{10, 20}
	arr3 := [5]int{10, 20, 30, 40, 45}

	arr3[4] = 50
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

}
