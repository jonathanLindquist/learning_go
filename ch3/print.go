package main

import "fmt"

const newConst = 10

func main() {
	var x = 1_234

	fmt.Printf("int: %t\n", x == 1234)

	var hex = 0xA6

	fmt.Printf("hex: %x\n", hex)

	var flag bool
	flag = true

	fmt.Printf("flag: %t\n", flag)

	var multiply int = 25
	multiply *= 5.0

	fmt.Printf("multiply: %d\n", multiply)

	var float float64 = 7.0
	float *= 3.7

	fmt.Printf("float: %f\n", float)

	fmt.Printf("converted float to int: %d\n", int(float))

	shortHand := 10
	shortHand = 15

	fmt.Printf("shorthand decl: %d\n", shortHand)

	fmt.Printf("const: %d\n", newConst)
}
