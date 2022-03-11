package main

import (
	"fmt"

	module "github.com/stheven26/module"
)

func main() {
	// fmt.Print(module.FeetToMeter(100))
	fmt.Print("Option:\n1.Convert Feet to Meter\n2.Convert Mile to KM\n3.Convert Liter to CC\n")
	fmt.Print("Enter Option Function: ")
	var Option int
	fmt.Scan(&Option)

	if Option == 1 {
		fmt.Print("Enter Feet: ")
		var feet float64
		fmt.Scan(&feet)
		meter := module.FeetToMeter(feet)
		fmt.Printf("Hasil Konversi %.2f Feet adalah %.2f Meter\n", feet, meter)
	} else if Option == 2 {
		fmt.Print("Enter Mile: ")
		var Mile float64
		fmt.Scan(&Mile)
		km := module.MileToKM(Mile)
		fmt.Printf("Hasil Konversi %.2f Mile adalah %.2f KM\n", Mile, km)
	} else if Option == 3 {
		fmt.Print("Enter CC: ")
		var CC float64
		fmt.Scan(&CC)
		liter := module.CCToLiter(CC)
		fmt.Printf("Hasil Konversi %.1f CC adalah %.2f Liter\n", CC, liter)
	}
}
