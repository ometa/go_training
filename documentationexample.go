// Package documentationexample is used as an example
package documentationexample

import "fmt"

// Baz is the only way to go when Foo and Bar are used.
//	This indent will be code.
const Baz = "Baz"


// Bar returns baz
func Bar() string {
	return Baz
}

// Prefixing a function name with Example will render as an example in godoc, and actually allow the code to be run.
func ExampleBar() {
	fmt.Println(Bar())
	// Output: Baz     // need the colon here.
}

func main() {
}
