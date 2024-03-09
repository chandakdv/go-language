package main

// import (
// 	"fmt"
// )

// func main() {
// 	var greet = func() {
// 		fmt.Println("hello,how are you")
// 	}
// 	var welcome = greet
// 	greet()
// 	welcome()
// }

//anonymous function with parameter
// 	var sum = func(n1 int, n2 int) {
// 		result := n1 + n2
// 		fmt.Println("sum is ", result)
// 	}
// 	sum(5, 3)

//return value from anonymous function
// 	var sum = func(n1, n2 int) int {
// 		sum := n1 + n2
// 		return sum
// 	}
// 	result := sum(5, 3)
// 	fmt.Println("sum is", result)

// anonymous function as argument to ither function
// var sum = 0
// func findSquare(num int) int {
// 	square := num * num
// 	return square
// }
// func main() {
// 	sum := func(number1 int, number2 int) int {
// 		return number1 + number2
// 	}
// 	result := findSquare(sum(6, 9))
// 	fmt.Println("result is", result)
// }

// // print odd numbers using golang
// func calculate() func() int {
// 	num := 1
// 	//returns inner functon
// 	return func() int {
// 		num = num + 2
// 		return num
// 	}
// }
// func main() {
// 	//call the outer function
// 	odd := calculate()
// 	//call the inner function
// 	fmt.Println(odd())
// 	fmt.Println(odd())
// 	fmt.Println(odd())

// 	//call the outer function again
// 	odd2 := calculate()
// 	fmt.Println(odd2())
// }
