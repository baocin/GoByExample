// PANIC
package main

import (
	"os"
)

func main() {
	panic("oh noes A problem has occured")

	_, error := os.Create("/tmp/file")
	if error != nil {
		panic(error)
	}
}
