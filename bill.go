package main

import (
	"fmt"
)

//create type for bill
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//create func to generate new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

//crate a func for format to call bill
func (b *bill) format() string {
	fs := fmt.Sprintf("Bill Breakdown for Sir/madam: %v \n", b.name)
	var total float64
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%0.2f \n", k+":", v)
		total += v
	}
	//creat tip
	fs += fmt.Sprintf("%-25v ...%0.2f \n", "Tip", b.tip)

	//calculate total amount value
	fs += fmt.Sprintf("%-25v ...%0.2f \n", "Total:", total+b.tip)
	return fs
}

//create func to update tip amount
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//create func to add new item to the bill menu
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price

}
