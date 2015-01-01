// arrays
package main

import (
	"fmt"
)

func main() {
	var a [5]int //array of 5 ints
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoDim [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2 dim: ", twoDim)
}
