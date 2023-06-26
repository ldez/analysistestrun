package foo

import "fmt"

// BUG: here // want "There is a bug"
func Foo() {
	fmt.Println("Foo")
}
