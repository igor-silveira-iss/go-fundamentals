package main

import (
	"fmt"
)

func main() {

	// ARRAY
	var array1 [5]string
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICE
	slice := []int{10, 20, 40, 50, 60, 78, 89}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(reflect.TypeOf(array3))

	// ARRAYS INTERNOS
	fmt.Println("\n--------------------------------\n")
	slice7 := make([]float32, 10, 11)
	slice7 = append(slice7, 5)
	slice7 = append(slice7, 6)
	fmt.Println(slice7)
	fmt.Println(len(slice7))
	fmt.Println(cap(slice7))

}
