package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ch4p1And2()
	ch4p3()
}

func ch4p1And2() {
	mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, rand.IntN(100))
	}
	fmt.Println(mySlice)

	// for i := 0; i < len(mySlice); i++ {
	// 	if mySlice[i]%3 == 0 && mySlice[i]%2 == 0 {
	// 		fmt.Println("Six!")
	// 	} else if mySlice[i]%2 == 0 {
	// 		fmt.Println("Two!")
	// 	} else if mySlice[i]%3 == 0 {
	// 		fmt.Println("Three!")
	// 	} else {
	// 		fmt.Println("Never mind")
	// 	}
	// }

	for _, value := range mySlice {
		// "_" is index but "_" is ignored and not tracked
		switch {
		case value%6 == 0:
			fmt.Println("Six!")
		case value%2 == 0:
			fmt.Println("Two!")
		case value%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func ch4p3() {
	total := 0
	for i := 0; i < 10; i++ {
		total := total + 1
		fmt.Println(total)
	}

	// Bug: Re-declaring total to 1 for each iteration of the for loop
	// Resulting in total equaling "1" for each iteration
}
