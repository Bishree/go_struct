package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// creat a function to get users input

func getInput(prompt string, r *bufio.Reader) (string, error) {
	//print something to the console
	fmt.Println(prompt)

	//get the input from console
	input, err := r.ReadString('\n')

	//return the output of the function
	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Create a new bill name:")
	//name, _ := reader.ReadString('\n')
	name, _ := getInput("Create a new bill name:", reader)
	fmt.Sprintf("The new bill has been created with the name: %v \n", name)
	b := newBill(name)
	return b
}
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills_"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("the bill saved")
}

func promptOptions(b bill) {
	//get the input from the console
	reader := bufio.NewReader(os.Stdin)

	//print the prompt to the console
	opt, _ := getInput("choose an option (a - add item, s - save bill, t - add tip)", reader)
	switch opt {
	case "a":
		name, _ := getInput("enter the item:", reader)
		price, _ := getInput("enter the price:", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("try again")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println(name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter the tip amount ($):", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("try again")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println(tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you bill saved", b.name)
		fmt.Println(b.format())
	default:
		fmt.Println("this is was not a valid option")
		promptOptions(b)

	}
}
