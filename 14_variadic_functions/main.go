package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// almost the same as above, with return value
func sum2(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {

	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	fmt.Print(nums, " ")
	fmt.Println(sum2(nums...))
}
