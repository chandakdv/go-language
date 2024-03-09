package main

// import (
// 	"fmt"
// 	"math"
//"unsafe"
//"reflect"
// )

// func main() {
// var a int = 5
// var b int = 3
// a := 5
// b := 3
// c := a & b
// d := a | b
// e := a ^ b

// fmt.Println(c)
// fmt.Println(d)
// fmt.Println(e)

//a := 10
// fmt.Println(a > 5)
// fmt.Println(a == 10)
// fmt.Println(a < 5)
// fmt.Println(a >> 5)

// var num1 int = 10
// var num2 byte = 20
// var num3 float64 = 10.20
// var num4 string = "Divya"
// fmt.Printf("size of num1:%d", unsafe.Sizeof(num1))
// fmt.Printf("\nSize of Num2 :%d", unsafe.Sizeof(num2))
// fmt.Printf("\n Size of num3 :%d", unsafe.Sizeof(num3))
// fmt.Printf("\n Size of num4:%d", unsafe.Sizeof(num4))

// a := 10
// b := 10.20
// c := "Hello"
// d := true
// e := []string{"India", "USA", "UK"}
// // printing the values
// fmt.Println("a = ", a)
// fmt.Println("b = ", b)
// fmt.Println("c=", c)
// fmt.Println("d=", d)
// fmt.Println("e=", e)
//printing the type of yhe variables
// fmt.Println("Type of a = ", reflect.TypeOf(a))
// fmt.Println("Type of b = ", reflect.TypeOf(b))
// fmt.Println("Type of c = ", reflect.TypeOf(c))
// fmt.Println("Type of d = ", reflect.TypeOf(d))
// fmt.Println("Type of e = ", reflect.TypeOf(e))
//printing the type of variables
// fmt.Println("Type of a = ", reflect.ValueOf(a).Kind())
// fmt.Println("Type of b = ", reflect.ValueOf(b).Kind())
// fmt.Println("Type of c = ", reflect.ValueOf(c).Kind())
// fmt.Println("Type of d = ", reflect.ValueOf(d).Kind())
// fmt.Println("Type of e = ", reflect.ValueOf(e).Kind())

//defer print all the statements which usese defer in reerse order
// defer fmt.Println("Hello")
// defer fmt.Println("hiii")
// defer fmt.Println("Good Morning")

//area of circle
// var radius float32
// var area float32
// fmt.Println("Enter Radius")
// fmt.Scanf("%g", &radius)
// area = 3.14 * radius * radius
// fmt.Printf("Area of Circle:%g", area)

// var radius float32
// var area float32
// fmt.Println("Enter Radius")
// fmt.Scanf("%f", &radius)
// area = 3.14 * radius * radius
// fmt.Printf("Area of Circle:%.2f", area)

// var radius float32
// var area float32
// const pi = 3.14
// fmt.Println("Enter Radius")
// fmt.Scanf("%g", &radius)
// area = pi * radius * radius
// fmt.Printf("Area of Circle:%g", area)

// var radius float32 = 0
// var perimeter float32 = 0
// fmt.Println("Enter Radius")
// fmt.Scanf("%g", &radius)
// perimeter = 2 * 3.14 * radius
// fmt.Println("Perimeter of a circle : %f", perimeter)

//converting the fahrenit to celsius in golang
// var ftemp float64 = 0
// var ctemp float64 = 0
// fmt.Printf("Enter temperature in fahrenheit:")
// fmt.Scanf("%f", &ftemp)
// ctemp = (ftemp - 32) / 1.8
// fmt.Printf("temperature in celsius:%f", ctemp)

// var ftemp float64
// var ctemp float64
// fmt.Printf("Enter temperature in celsius:")
// fmt.Scanf("%f", &ctemp)
// ftemp = (ctemp * 1.8) + 32
// fmt.Printf("temperature in  fahrenheit:%.2f", ftemp)

//squareroot of number
// 	var x int = 225
// 	var r float64
// 	r = math.Sqrt(float64(x))
// 	fmt.Printf("Square root of %d is %.2f\n", x, r)
// }
