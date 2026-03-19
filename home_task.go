package main

import "fmt"

func main() {
	//Home task 1
	// fmt.Println("Easy:")
	// easy1()
	// fmt.Println("Middle:")
	// middle1()
	// fmt.Println("Hard:")
	// hard1()

	// //Home task 2
	// fmt.Println("Middle:")

	// moto := Moto{
	// 	Brand:  "Bajaj",
	// 	Model:  "Pulsar180",
	// 	Color:  "red",
	// 	Number: "С567ТР",
	// }

	// car := Car{
	// 	Brand:  "Geely",
	// 	Model:  "Coolray",
	// 	Color:  "pink",
	// 	Number: "А234ПС",
	// }

	// StartRide(&moto)
	// StartRide(&car)

	// StopRide(&moto)
	// StopRide(&car)

	fmt.Println("Hard:")

	data := []byte("more text")

	file := File{
		Name: "hello.txt",
	}

	cons := Console{}

	i, err := WriteData(&file, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Записано", i, "байт в файл")
	}

	i, err = WriteData(&cons, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Записано", i, "байт в консоль")
	}
}
