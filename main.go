package main

import (
	"fmt"

	"github.com/EarthlyZ9/learngo/somefunc"
)

func main() {
	fmt.Println("Hello world!")
	// Only Upper Case functions are exported
	somefunc.SayHello()

	const name = "earthlyz9"              // Untyped variable
	const typed_name string = "earthlyz9" // Typed variable

	// typed_name = "earth" // cannot modify constant variable

	var var_name string = "jisoo"
	var_name = "js"
	fmt.Println(var_name)

	// short hand variable declaration (inside function only)
	short_name := "jisoo" // Go guesses the type automatically
	short_name = "js"
	fmt.Println(short_name)

}
