package main

import "fmt"

func main() {

	fmt.Println("Magic of array")

	// Array

	// var ages [3]int = [3]int{45, 56, 22}
	var ages = [3]int{45, 56, 22}

	fmt.Println(ages, len(ages))

	names := [3]string{"Abby", "Vivek", "Amit"}
	names[0] = "Abishek"

	fmt.Println(names, len(names))

	// Slices

	var scores = []int{50, 45, 23}

	scores[1] = 12

	scores = append(scores, 56)

	fmt.Println(scores, len(scores))

	// slice rages

	rangeOne := names[0:2]

	rangerTwo := names[0:]

	rangeThree := names[:1]

	fmt.Println(rangeOne)
	fmt.Println(rangerTwo)
	fmt.Println(rangeThree)

	// appending to slices

	rangerTwo = append(rangerTwo, "Rani")
	fmt.Println(rangerTwo)

}
