package main

import "fmt"

func main() {
	// assigning array of any length
	grades := [...]int{97, 94, 93}
	fmt.Printf("Grades: %v\n", grades)

	// creating array with a length of 3
	var students [3]string
	fmt.Printf("Students: %v\n", students)

	// assigning values to indexes
	students[0] = "Matthew"
	students[1] = "Michael"
	students[2] = "Brian"
	fmt.Printf("Students: %v\n", students)

	// returning values at a given index
	fmt.Printf("Student #2: %v\n", students[2])

	// length of Array
	fmt.Printf("Number of Students: %v\n", len(students))

	// array of arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	//When you make a copy of an array it makes an exact copy and if using a ton of elements can impact performance
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
