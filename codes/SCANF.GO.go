package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	fmt.Printf("Enter name and age ")
	if _, err := fmt.Scan(&name, &age); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Your name is %s\n", name)
	fmt.Printf("Your age is %d\n", age)
	fmt.Printf("Your name is %s\n", name)
}
