package main

import "fmt"

func main() {

	//TODO create a bill structure inside bill file
	//call this bill here

	mybill := newBill("ahmed Yaseen")
	mybill.updateTip(50.5254)
	mybill.addItem("anber rise", 55.7584)
	mybill.addItem("egg", 25.365)

	fmt.Println(mybill.format())
	//type bill
}
