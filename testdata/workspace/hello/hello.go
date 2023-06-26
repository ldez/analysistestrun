package main

import (
	"fmt"

	"example.com/hello/foo"
	"golang.org/x/example/stringutil"
)

// BUG: here // want "There is a bug"
func main() {
	foo.Foo()
	fmt.Println(stringutil.ToUpper("Hello"))
}
