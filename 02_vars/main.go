package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Ganesh"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Ganesh"
	// email := "ganesh@mail.com"

	name, email := "Ganesh", "ganesh@mail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
