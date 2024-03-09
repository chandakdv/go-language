//golang program to illustrate the usage of fmt.Printf() function

// Including the main Package
package main

//Import fmt
import (
	"fmt"
)

// Calling main
func main() {

	// const name, dept = "GreeksforGreeks", "CS"

	// fmt.Printf("%s is a %s Portal.\n", name, dept)

	//var a int =10
	//var b float=10.6
	// var name string="divya"
	//fmt.Printf("%d %f %s",a,b,name)

	var str = "GreeksforGreeks"
	fmt.Printf("The string is %s \n", str)
	var num1 int = 21
	fmt.Printf("the decimal valye is %d \n", num1)
	var num2 float32 = 7.786
	fmt.Printf("the floating point is %g \n", num2)
	var num3 int = 1
	fmt.Printf("the binary value of num3 is %b \n", num3)
	var num4 = 4 + 11
	fmt.Printf("Scientific notation of num 4: %e \n", num4)

}
