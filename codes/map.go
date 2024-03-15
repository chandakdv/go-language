package main

// import (
// 	"fmt"
// 	//"sort"
// )

// func main() {
// 	//create a map
// 	subjectMarks := map[string]float32{"Golang": 45, "Java:": 40, "Python:": 81}
// 	fmt.Println(subjectMarks)
// }

// // access values of a map in golang
// func main() {
// 	//create a map
// 	flowercolor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}
// 	//access valu for key sunflower
// 	fmt.Println(flowercolor["Sunflower"])
// }

// //change value of a map
// func main() {
// 	//create a map
// 	capital := map[string]string{"Nepal": "Kathmandu", "US": "NewYork"}
// 	fmt.Println("Initial Map", capital)
// 	//change value of US to Wasington DC
// 	capital["US"] = "Wahington DC"
// 	fmt.Println("Udated Map:", capital)
// }

// func main() {
// 	//create a map
// 	students := map[int]string{1: "john", 2: "lilly"}
// 	fmt.Println("initial map", students)
// 	//add elements with key 3
// 	students[3] = "Robin"
// 	//add elements with key 5
// 	students[5] = "Julle"
// 	fmt.Println("updated value:", students)
// }

// // delete elements of go map elements
// func main() {
// 	//create a map
// 	personage := map[string]int{"Hemione": 21, "Harry": 20, "John": 25}
// 	fmt.Println("Initial map:", personage)
// 	//remove element of map with key john
// 	delete(personage, "John")
// 	fmt.Println("updates map:", personage)
// }

// // looping through map
// func main() {
// 	//create a map
// 	squarednumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}
// 	for number, squared := range squarednumber {
// 		fmt.Printf("square of %d is %d\n", number, squared)
// 	}
// }

// func main() {
// 	//create a map using make
// 	student := make(map[int]string)
// 	//add elements to the map
// 	student[1] = "Harry"
// 	student[2] = "Lilly"
// 	student[3] = "Harminoie"
// 	fmt.Println(student)
// }

// func main() {
// 	//create a map
// 	places := map[string]string{"nepal": "ksttmzndu", "US": "Wadington DC", "Norway": "Oslo"}
// 	//access only the keys of the map
// 	for country := range places {
// 		fmt.Println(country)
// 	}
// }

// //map of using structure
// type Vertex struct {
// 	Lat, Long float64
// }
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }
// func main() {
// 	fmt.Println(m)
// 	m["Nagpur"] = Vertex{379.293, -287.868}
// 	fmt.Println(m)
// }

// func main() {
// 	places := map[string]string{"nepal": "kattmzndu", "US": "Wadington DC", "Norway": "Oslo"}
// 	for _, capital := range places {
// 		fmt.Println(capital)
// 	}
// }
