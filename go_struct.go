package main

import "fmt"

func main() {

	//TODO create a bill structure inside bill file
	//call this bill here

	mybill := newBill("ahmed Yaseen")
	mybill.updateTip(50)
	mybill.addItem("anber rise", 55)
	mybill.addItem("egg", 25)
	mybill.addItem("orange jus", 12)
	mybill.addItem("chicken", 48)
	mybill.addItem("Texas steak", 120)
	mybill.addItem("coffee espresso", 10)
	fmt.Println(mybill.format())
	//type bill
}
