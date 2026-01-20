package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = max(0, 18-age)

	yearsUntilDrinking = max(0, 21-age)

	yearsUntilCarRental = max(0, 25-age)

	// naked return
	return
}
func main() {
	adult, drinking, car := yearsUntilEvents(13)
	fmt.Println(adult, drinking, car)

}
