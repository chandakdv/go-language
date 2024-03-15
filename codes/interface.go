//example of a go program that makes an http get request to an API

//go run program_name.go

//go build program_name.go
//prohram_name.exe



package main
import (
	"fmt" 
"math"
)

//define an interface
type Shape interface{
	area() float64
}

//define a circle
type Circle struct{
	x,y,radius float64
}

//define a rectangle
type Rectangle struct{
	width , height float64
}

//define a methode for circle (impelemntation of Shape.area())
func (circle Circle) area() float64{
	return math.Pi*circle.radius*circle.radius
}

//define a methode for rectangle (implementation of Shape.area())
func (rect Rectangle) area() float64{
	return rect.width*rect.height
}

//define a methode for shape
func getArea(shape Shape)float64{
	return shape.area()
}

func main(){
	circle:=Circle(x:0,y:0,radius:5)
	rectangle:=Rectangle(width:10,height:5)

	fmt.Printf("Circle area:%f\n",getArea(circle))
	fmt.Printf("Rectangle area:%f\n",getArea(rectangle))
}