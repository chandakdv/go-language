//write a program to print largest among three number

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	fmt.Printf("Enter number 1 ")
	fmt.Scan(&num1)
	fmt.Printf("Enter number 2 ")
	fmt.Scan(&num2)
	fmt.Printf("Enter number 3 ")
	fmt.Scan(&num3)
	if num1 > num2 && num1 > num3 {
		fmt.Println("Num 1 is biggest")
	}
	if num2 > num1 && num2 > num3 {
		fmt.Println("Num 2 is biggest")
	}
	if num3 > num2 && num3 > num1 {
		fmt.Println("Num 1 is biggest")
	}
}
