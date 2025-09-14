package main

import (
	"errors"
	"fmt"
)

func main() {
	var name1 string = "Humayoon"
	var age1 int = 16
	var name, age, err = Details(name1, age1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Welcome to HSX Mr. %v. You're %v years old\n", name, age)
	}
}

func Details(Name1 string, age int) (string, int, error) {
	var err error
	if Name1 != "Humayoon" || age < 18 {
		return "", 0, errors.New("Intruder")

	}
	return Name1, age, err
}
