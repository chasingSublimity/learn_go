package main

import "fmt"

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// }

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83
//
// 	var total float64
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	fmt.Println(total / float64(len(x)))
// }

// func main() {
// 	x := [5]float64{
// 		98,
// 		93,
// 		77,
// 		82,
// 		83,
// 	}
// 	var total float64
// 	for _, value := range x {
// 		total += value
// 	}
// 	fmt.Println(total / float64(len(x)))
// }

// slices
// func main() {
// 	// var x []float64
// 	// x := make([]float64, 5)
// 	// x := make([]float64, 5, 10)
// 	arr := [5]float64{1, 2, 3, 4, 5}
//             // High is exclusivek
// 	x := arr[0:5]
//
// 	fmt.Println(x)
// }

// slice append
// func main() {
// 	slice1 := []int{1, 2, 3}
// 	/* if append() causes a slice to exceed
// 	the current array length, a new slice is created
// 	and returned */
// 	slice2 := append(slice1, 4, 5)
// 	fmt.Println(slice1, slice2)
// }

// slice copy
// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := make([]int, 2)
// 	copy(slice2, slice1)
// 	fmt.Println(slice1, slice2)
// }

// maps
// func main() {
// 	x := make(map[string]int)
// 	x["key"] = 10
// 	fmt.Println(x["key"])
// }

// func main() {
// 	x := make(map[int]int)
// 	x[1] = 10
// 	fmt.Println(x[1])
// }

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	// fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
}
