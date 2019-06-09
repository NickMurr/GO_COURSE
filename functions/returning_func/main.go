// functions(14)(10) package main

package main

import "fmt"

func main() {
	// x := bar()
	// fmt.Println(x())

	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}
