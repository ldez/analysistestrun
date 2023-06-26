package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// BUG: here // want "There is a bug"
func main() {
	fmt.Println(stringutil.ToUpper("World"))
}
