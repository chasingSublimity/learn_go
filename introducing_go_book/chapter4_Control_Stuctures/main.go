package main

import "fmt"

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "even")
// 		} else {
// 			fmt.Println(i, "odd")
// 		}
// 	}
// }

// func main() {
// 	var i int
// 	switch i {
// 	case 0:
// 		fmt.Println("zero")
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	case 4:
// 		fmt.Println("four")
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
