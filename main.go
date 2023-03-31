package main

import (
	"fmt"
	"strings"

	"github.com/EarthlyZ9/learngo/somefunc"
)

func multiply(a, b int) int {
	return a * b
}

// Go function can return multiple values
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// initalizing variables when declaring return types
func lenAndUpper2(name string) (length int, uppercase string) {
	// do something after return
	defer fmt.Println("done!")

	length = len(name)
	uppercase = strings.ToUpper(name)

	// return length, uppercase
	return // naked return
}

// multiple arguments
func repeatMe(words ...string) {
	fmt.Println(words)
}

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

	// using functions
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("jisoo")
	// totalLength, _ := lenAndUpper("jisoo")

	fmt.Println(totalLength, upperName)

	repeatMe("js", "earthlyz9", "jisoo", "leejisoo") // returns array

	fmt.Println(lenAndUpper2("jisoo"))

}
