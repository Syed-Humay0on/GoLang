package main

import (
	"fmt"
	_ "os" // add a underscore prefix to get rid of compile time error for unused packages
)

func main() {
	// a curly bracket { cannot appear in the begining of a line in go
	// comments start with double slash or slash aestrik for multi line code
	/*
		multi
		line
		comment
	*/
	myName := "Emma Watson" //decalre a variable using := (only inside a function)
	myName = "Zain Malik"   // reassign that declared variable

	//switch case
	switch myName {

	case ("Emma Watson"):
		fmt.Println("Queen")

	case ("Zain Malik"):
		fmt.Println("Chad")

	default:
		fmt.Println("Absolutely Nobody")
	}
}
