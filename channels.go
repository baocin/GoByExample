// channels
package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	go func() { mesage <- "ping" }()
	msg := <-message
	fmt.println(msg)
}
