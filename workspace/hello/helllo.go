package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}

/*
create int.go in workspace/example/hello/reverse
and put this code
package reverse

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}
*/
