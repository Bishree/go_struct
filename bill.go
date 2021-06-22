package main

import (
	"fmt"
	"strings"
)

//creat bill type structure
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//creat new bill form bill type
func newBill(name string) bill {
	b := bill{
		name:  strings.ToUpper(name),
		items: map[string]float64{"Texas steak": 120, "anber rise": 55, "coffee espresso": 10},
		tip:   10,
	}
	return b

}

//creat a bill format receiver func
func (b bill) format() string {
	fs := fmt.Sprintf("Bill breakdown for %v is : \n", b.name)
	var total float64 = 0
	// to go through all items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%v \n", k+":", v)
		total += v
	}

	total += b.tip
	// getting total
	fs += fmt.Sprintf("%-25v ...%0.2f \n", "Total is:", total)

	// adding tip to the bill
	fs += fmt.Sprintf("%-25v ...%v", "Tip is:", b.tip)

	//return the bill format
	return fs
}
