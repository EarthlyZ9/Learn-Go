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

func superAdd(numbers ...int) int {
	// for loop
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	// ignore index by using _

	return 1
}

func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func canIDrink(age int) bool {
	// if else
	// creating variable inside if clause - variable expression
	// "created koreanAge var only for the if clause"
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
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

	superAdd(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(canIDrink(16))

	// pointers
	a := 2
	b := a // copy of value
	a = 10
	fmt.Println(a, b)   // 10, 2
	fmt.Println(&a, &b) // different memory address

	a = 2
	c := &a            // c is pointing at the address of a
	fmt.Println(&a, c) // same address value
	fmt.Println(*c)    // value of the address that c is pointing at -> 2

	*c = 20        // updating value of a using c
	fmt.Println(a) // -> 20

	// arrays

	names := [7]string{"js", "jisoo", "eartlyz9", "z9", "leejisoo"} // constant space (length)
	names[5] = "JS"
	names[6] = "LJS"

	// slice
	infinite_names := []string{"js", "jisoo", "eartlyz9", "z9", "leejisoo"}
	fmt.Println(infinite_names)
	infinite_names = append(infinite_names, "LJS") // append function alone does not modify the array itself, it just returns a new slice
	fmt.Println(infinite_names)

	// map
	js := map[string]string{"name": "jisoo", "age": "24"} // type is restricted to 'string'
	fmt.Println(js)
	for key, value := range js {
		fmt.Println(key, value)
	}

	// struct
	// Go struct does not have constructor method like __init__ or constructor()
	type person struct {
		name    string
		age     int
		favFood []string
	}

	favFood := []string{"icecream", "coffee"}
	// z9 := person{"jisoo", 24, favFood}
	z9 := person{name: "jisoo", age: 24, favFood: favFood}
	fmt.Println(z9)

}
