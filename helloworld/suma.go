package main

import "fmt"

func Suma() {
	fmt.Println("Hello world")

	var num1, num2 int

	fmt.Println("Numero 1: ")
	fmt.Scanln(&num1)

	fmt.Println("Numero 2: ")
	fmt.Scanln(&num2)

	suma := num1 + num2

	fmt.Printf("El resultado es: %d", suma)
}
