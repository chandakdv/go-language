package main

// import (
// 	"fmt"
// 	//"sort"
// )

// func main() {
// 	//create two slices
// 	primeNumbers := []int{2, 3, 4, 7}
// 	numbers := []int{1, 2, 3}
// 	//copy elements of prime numbers to numbers
// 	copy(numbers, primeNumbers)
// 	//print numbers
// 	fmt.Println("Numbers:", numbers)
// }

// //looping through go slice
// func main() {
// 	numbers := []int{2, 4, 6, 8, 10}
// 	//for loop that iterates through the slice
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }

// //iterating a slice using 'range' in 'for' loop
// func main() {
// 	//create an ineteger array
// 	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
// 	//create slice of from index 2 till index4(5,-1)
// 	inSlice := arr[2:5]
// 	fmt.Println("slice elements:")
// 	for index, ele := range inSlice {
// 		fmt.Printf("index=%d,elements=%d\n", index, ele)
// 	}
// }

// func main() {
// 	slice := []string{"honesty", "is", "the", "best", "policy"}
// 	sort.Strings(slice)
// 	fmt.Println("Sorted slice:")
// 	for _, item := range slice {
// 		fmt.Printf("%s ", item)
// 	}
// }
