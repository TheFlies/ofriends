package strlib_test

import (
	"fmt"

	strlib "github.com/nasa9084/go-strlib"
)

// ExampleSnakeCase is an example of usage of SnakeCase(string) function
func ExampleSnakeCase() {
	const camel = "CamelCaseString"
	fmt.Printf("%s\n", strlib.SnakeCase(camel))
}
