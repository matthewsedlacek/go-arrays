package main

import "fmt"

func main() {
	//slice
	a := []int{1, 2, 3}
	fmt.Println(a)

	// Capacity and length can be different. A Slice can only be a portion of the actual array
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
