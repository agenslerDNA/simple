package main

import "fmt"

func main() {
	fmt.Println("Starting")
	fmt.Printf("1 + 1 = %d\n", Add(1, 1))
	fmt.Println("Done")
}

func Add(l, r int) int {
	return l + r
}
