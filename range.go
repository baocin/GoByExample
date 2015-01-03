// range
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sums := 0
	for _, num := range nums {
		sums += num
	}
	fmt.Println("sums: ", sums)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i, "  value: ", num)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
