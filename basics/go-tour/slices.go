package main

import "fmt"

// slice create
func sliceConstructors() {
	fmt.Println("**********CREATING SLICES**********")
	//1. From Existing Array
	arr := [5]int{1, 4, 9, 16, 25}
	slice1 := arr[0:3]
	fmt.Println("Slice From Array", slice1)

	//2.Literal
	slice2 := []string{"C", "C++", "Java", "Go", "Python"} // Capacity== Length = 5
	fmt.Println("Slice Literal", slice2)

	//3.From another slice
	slice3 := slice2[:2]
	fmt.Println("Slice From Another Slice", slice3)

	// 4. make function make(Type, Length, Capacity)
	slice4 := make([]float32, 3) // Capacity== Length
	fmt.Println("Slice using make function", slice4)
}

func sliceIteration() {
	fmt.Println("**********ITERATING SLICES**********")
	// loop over slices
	myslice := []int{1, 4, 9, 16, 25}
	fmt.Println("Iteratin over slice start")
	for i, v := range myslice {
		fmt.Println("myslice[", i, "] =", v)
	}
}

func nilSlices() {
	fmt.Println("**********NIL SLICES**********")
	var nilSlice []int // nil slice
	var emptySlice []int = []int{}
	fmt.Println("nilSlice is Nil ?", nilSlice == nil)
	fmt.Println("emptySlice is Nil ?", emptySlice == nil)
}

func lenAndCapacity() {
	fmt.Println("**********Length & Capacity**********")
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[1:4]
	fmt.Println("Array ", arr)
	// len =  high-low = 4-1 = 3
	// cap = arr_len - slice start index = 6 - 1 = 5
	fmt.Println("Slice ", slice, "Length is", len(slice), "Capacity is", cap(slice))
}

func sliceDereference() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	slice := arr[1:4]

	fmt.Println("****************************************")
	fmt.Println("Original Array ", arr)
	fmt.Println("Original Slice ", slice)
	fmt.Println("Slice Length is", len(slice), "Capacity is", cap(slice))

	slice = append(slice, 9)
	fmt.Println("--------------")
	fmt.Println("After inserting 4th element in slice ", "Length is", len(slice), "Capacity is", cap(slice))
	// adding an element to slice, array got updated
	fmt.Println("Updated arr is", arr)

	slice = append(slice, 99)
	fmt.Println("--------------")
	fmt.Println("After inserting 5th element in slice ", "Length is", len(slice), "Capacity is", cap(slice))
	fmt.Println("Updated arr is", arr)

	slice = append(slice, 999)
	fmt.Println("--------------")
	fmt.Println("After inserting 6th element in slice ", "Length is", len(slice), "Capacity is", cap(slice))
	fmt.Println("Updated arr is", arr)

	slice = append(slice, 9999)
	fmt.Println("--------------")
	fmt.Println("Now Slice reached its max capacity, it will not mutate the array now, It will point to newly allocated array")
	fmt.Println("After inserting 7th element in slice ", "Length is", len(slice), "Capacity is", cap(slice))
	fmt.Println("Updated arr is", arr)
}

func main() {
	sliceConstructors()
	nilSlices()
	sliceIteration()
	lenAndCapacity()
	sliceDereference()
}
