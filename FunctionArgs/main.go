/*
	This is a test experiment for function args.
	In golang, function will copy a new variable for args, and it calls by value.
	Modification of will not affect the variable we input. (They are totally distinct.)

	If we want to modify the struct in input parameters by reference,
	we have to use pointer to point to the address in Heap memory.
*/

package main

import (
	"fmt"
)

type Object struct {
	key string
}

func modifyPrimitive[T int64 | float64](x T) {
	x = 100
}

func modifyList[T [3]int | []int](a T) {
	a[0] = 100
}

func modifySlice(a []int) {
	a[0] = 100
}

func modifyMap(m map[string]string, p *map[string]string) {
	m["A"] = "Modified"
}

func modifyStruct(o Object) {
	o.key = "Modified"
}

func modifyStructPointer(o *Object) {
	o.key = "Modified"
}

func main() {
	// Test for primitive type
	var primitiveValue int64 = 30

	modifyPrimitive(primitiveValue)
	fmt.Printf("Primitive value after modify: %v\n", primitiveValue) // 30 (no change)

	// Array
	var testArray = [3]int{1, 2, 3}
	modifyList(testArray)
	fmt.Printf("Array value after modify: %v\n", testArray) //[1, 2, 3] (no change)

	// Slice
	/*
		In fact slice is made by a struct, which store { pointer, len, cap }
		As a result, when sending slice into func, we use struct pointer to modify value.
	*/
	var testSlice = []int{1, 2, 3}
	modifyList(testSlice)
	fmt.Printf("Slice value after modify: %v\n", testSlice) //[100, 2, 3] (changed)

	// Map
	testMap := map[string]string{
		"A": "hello",
		"B": "world",
		"C": "!!",
	}

	modifyMap(testMap, &testMap)
	fmt.Printf("Map value after modify: %v\n", testMap) // map[A:Modified B:world C:!!]  (changed)
	
	// Struct
	object := Object {
		key: "Hello world",
	}

	modifyStruct(object)
	fmt.Printf("Struct value after modify: %v\n", object) // {Hello world} (no change)

	modifyStructPointer(&object)
	fmt.Printf("Struct value after modify: %v\n", object) // {Modified} (changed)
}
