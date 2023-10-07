package main

import "fmt"

func main() {

	// While type lop

	x := 0

	for x < 5 {
		fmt.Printf("The vaue of x is %v\n", x)
		x++
	}

	// For loop

	for i := 0; i < 5; i++ {
		fmt.Printf("the value of i is %v\n", i)
	}

	// slice of names
	names := []string{"Abishek", "Vivek", "Amit", "Rani"}

	// printing a each value from slice till the end of it/
	for j := 0; j < len(names); j++ {
		fmt.Println(names[j])
	}

	// printing the value and index from a slice
	for index, value := range names {
		fmt.Printf("The index at %v has value %v \n", index, value)
	}

	// wht if you ant only to pring the value and not index
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

}
