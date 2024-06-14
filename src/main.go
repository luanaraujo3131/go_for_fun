package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func sub(x float64, y float64) float64 {
	return x - y
}

// Functions in go can result 2 values, in this case, a float64
// and a boolean
func mult(x float64, y float64) (float64, bool) {
	if x > 100 {
		return x * y, true
	} else {
		return x * y, false
	}
}

func main() {
	/*
		var name string
		name = "Luan"
	*/

	name := "Luan"
	fmt.Println("Name: ", name)

	fmt.Println(sum(300, 450))
	fmt.Printf("%.2f\n", sub(356.6, 180.9))

	//mult
	//float bool
	result, status := mult(150, 2)
	fmt.Println("Result:", result, ", status:", status)

}
