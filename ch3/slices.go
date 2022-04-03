package main

import "fmt"

func main() {
	var x = [3]int{10, 20, 30}

	fmt.Printf("int array: %o\n", x)

	var y = []int{10, 20, 30}

	fmt.Println(y)
	y = append(y, 40)
	fmt.Println(y)
	fmt.Println(cap(y))
}
