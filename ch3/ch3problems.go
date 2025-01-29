package main

import "fmt"

func ch3a1() {
	var greetings = [5]string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	var s1 = greetings[:2]

	var s2 = greetings[1:4]

	var s3 = greetings[3:]

	fmt.Println(greetings)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func ch3a2() {
	var message string = "Hi 👩 and 👨"

	var x = []rune(message)

	fmt.Println(string(x[3]))
}

func ch3a3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	joe := Employee{
		"john",
		"doe",
		25,
	}

	vansh := Employee{
		firstName: "vansh",
		lastName:  "kanojia",
		id:        22,
	}

	var bill Employee

	bill.firstName = "bill"
	bill.lastName = "gates"
	bill.id = 2

	fmt.Println(joe)
	fmt.Println(vansh)
	fmt.Println(bill)
}
