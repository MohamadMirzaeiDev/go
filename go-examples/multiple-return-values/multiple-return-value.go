package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Print(a)
	fmt.Print(b)


	_ , c := vals()

	fmt.Print(c)
}