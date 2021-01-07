package main

import "fmt"

func main() {
	a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// Adding an element to a slice
	a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// Concatenating Slices
	a = append(a, []int{2, 3, 4, 5}...)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// Stack Operations

	// remove element from the beginning
	// b := a[1:]
	// fmt.Println(b)

	// remove element from the end of an array
	// b := a[:len(a)-1]
	// fmt.Println(b)

	// remove element from the middle of an array
	b := append(a[:2], a[3:]...)
	fmt.Println(b)

}
