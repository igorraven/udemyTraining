package main

import "fmt"

func main() {
	var (num1, num2 int
	input float32)

	num1 = 3
	num2 = 9


	fmt.Println("\n\t Hello World From Jet Brain!")
	fmt.Printf("\n\t ( Num1= %d  Num2= %d ) = %d", num1, num2, num1+num2)

	fmt.Println("\n\t Hello World from Eclipse NEON!")
	fmt.Printf("\n\t ( Num1= %d  Num2= %d ) = %d", num1, num2, num1+num2)

	// An Introduction To Programming GO  NooK page 17 / 48
	fmt.Print("\n\t Enter a number: ")
	fmt.Scanf("%f", &input)
	output := (input *2)
	fmt.Println("\n\t Input squared = ", output)
}
