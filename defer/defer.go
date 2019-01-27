package main

import "fmt"

func main() {
	defer fmt.Println("5 world")

	fmt.Println("1 hello")

	fmt.Println("2 counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println("4 ", i)
	}

	fmt.Println("3 done")
}