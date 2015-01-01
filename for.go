// for
package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		//if this is an empty line then liteide auto deletes the next line
		fmt.Println("infinite loop :(")
		break
	}
}
