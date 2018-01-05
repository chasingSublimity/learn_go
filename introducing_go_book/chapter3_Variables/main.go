package main

import "fmt"

// func main() {
// 	var x string = "Hello, World"
// 	fmt.Println(x)
// }

// func main() {
// 	var x string
// 	x = "Hello, World"
// 	fmt.Println(x)
// }

// func main() {
// 	var x string
// 	x = "first"
// 	fmt.Println(x)
// 	x = "second"
// 	fmt.Println(x)
// }

// func main() {
// 	var x string
// 	x = "first"
// 	fmt.Println(x)
// 	x += "second"
// 	fmt.Println(x)
// }

// func main() {
// 	var x string = "hello"
// 	var y string = "world"
// 	fmt.Println(x == y)
// }

// func main() {
// 	var x string = "hello"
// 	var y string = "hello"
// 	fmt.Println(x == y)
// }

// func main() {
// 	x := "Hello, "
// 	var y = "World"
// 	z := 5
// 	fmt.Println(x + y)
// }

// defining multiple variables at once
// var (
// 	a = 5
// 	b = 10
// 	c = 15
// )

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
