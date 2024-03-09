package main

// import (
// 	"fmt"
// )

// closure is a nested function that helps to access the outer function's variable even after the outer function is closed
//outer function
// func greet() func() string {
// 	//variable definrd outside the inner function
// 	name := "john"
// 	//return a nested anonymous function
// 	return func() string {
// 		name = "hi" + name
// 		return name
// 	}
// }
// func main() {
// 	//call the outer function
// 	message := greet()
// 	//call the inner function
// 	fmt.Println(message())
// }

// closure helps in Data Isolation
// func displayNumbers() func() int {
// 	number := 0
// 	//inner function
// 	return func() int {
// 		number++
// 		return number
// 	}
// }
// func main() {
// 	//returns a closure
// 	//num1 := displayNumbers()
// 	fmt.Println(displayNumbers())
// 	fmt.Println(displayNumbers())
// 	//return a new closure
// 	//num2 := displayNumbers()
// 	fmt.Println(displayNumbers()())
// 	fmt.Println(displayNumbers()())
// }

// // fibonacci series
// func main() {
// 	var fib func(n int) int
// 	fib = func(n int) int {
// 		if n < 2 {
// 			return n
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}
// 	fmt.Println(fib(7))
// }
