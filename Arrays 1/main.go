package main
/*Adds elements to an array*/
/*Sorts the array*/
import (
	"fmt"
	"sort"
)
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSliceSorted(s []int) {
	fmt.Print("The Sorted Array ", s)
}
func main() {
	var s []int
	printSlice(s)
	// append works on nil slices.
	s = append(s, 4)
	printSlice(s)
	// The slice grows as needed.
	s = append(s, 3)
	printSlice(s)
	// We can add more than one element at a time.
	s = append(s, 2, 1, 0)
	printSlice(s)
	// Reordering Elements
	sort.Ints(s);
	printSliceSorted(s);
}