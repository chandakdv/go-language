package main

// import (
// 	"fmt"
// 	//"time"
// )

//program of a leap year
// func main() {
// 	var year int = 0
// 	fmt.Printf("enter number:")
// 	fmt.Scanf("%d", &year)
// 	if year != 0 {
// 		if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%400 == 0) {
// 			fmt.Printf("entered year is leap year ")
// 		} else {
// 			fmt.Printf("entered year is not a leap year ")
// 		}
// 	} else {
// 		fmt.Printf("enter the coreect number")
// 	}
// }

//nested if
// func main() {
// 	x := 8
// 	y := 10

// 	if x != y {
// 		if x < y {
// 			fmt.Println("x is less than y ")
// 		} else if x > y {
// 			fmt.Println("x is greater than y")
// 		}
// 	} else {
// 		fmt.Println("x is equal toy")
// 	}
// }

// func main() {
// 	var time float64 = 0.0
// 	fmt.Printf("enter the time:")
// 	fmt.Scanf("%f", &time)
// 	if time <= 23.59 {
// 		if time <= 10 {
// 			fmt.Printf("Good Morning")
// 		} else if time > 10 && time <= 20 {
// 			fmt.Printf("good afternoon")
// 		} else {
// 			fmt.Printf("Good evening")
// 		}
// 	} else {
// 		fmt.Printf("enter the correct time")
// 	}
// }

//switch statements
// func main() {
// 	thismonth := 5
// 	switch thismonth {
// 	case 1:
// 		fmt.Println("January")
// 	case 2:
// 		fmt.Println("Feb")
// 	case 3:
// 		fmt.Println("march")
// 	case 4:
// 		fmt.Println("april")
// 	case 5:
// 		fmt.Println("may")
// 	case 6:
// 		fmt.Println("june")
// 	}
// }

// func main() {
// 	today := time.Now()
// 	switch today.Day() {
// 	case 5:
// 		fmt.Println("today is fifth.clean your house")
// 	case 10:
// 		fmt.Println("today is 10th.buy some wine")
// 	case 15:
// 		fmt.Println("today is 15th . visit a doctor")
// 	case 25:
// 		fmt.Println("today is 15th.Buy some food")
// 	case 31:
// 		fmt.Println("today is 31.party tonight")
// 	default:
// 		fmt.Println("no info available for the day")
// 	}
// }

// func main() {
// 	x := 1

// 	switch x {
// 	case 1:
// 		fmt.Println("1")
// 		fallthrough
// 	case 3:
// 		fmt.Println("3")
// 		//fallthrough
// 	case 5:
// 		fmt.Println("5")
// 	}
// }

// func main() {
// 	//var x interface{}
// 	var x interface{} = "RKNEC"
// 	switch i := x.(type) {
// 	case nil:
// 		fmt.Printf("type of x:%T", i)
// 	case int:
// 		fmt.Printf("x is int")
// 	case float64:
// 		fmt.Printf("x is float")
// 	case func(int) float64:
// 		fmt.Printf("x is func(int)")
// 	case bool, string:
// 		fmt.Printf("x is bool or string")
// 	default:
// 		fmt.Printf("don't know the type")
// 	}
// }
