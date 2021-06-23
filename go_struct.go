package main

import "fmt"

func main() {
	//mybill := createBill()
	//promptOptions(mybill)

	//fmt.Println(mybill.format())
	//type bill

	shapes := []shape{
		square{length: 15.2},
		square{length: 4.9},
		circle{radius: 7.5},
		circle{radius: 12.3},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("___")
	}

}
