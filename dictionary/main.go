package main

import (
	"fmt"
	"log"

	"github.com/EarthlyZ9/learngo/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "firstWord"}
	definition, err := dictionary.Search("first")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(definition)

	dictionary.AddWord("second", "secondWord")
	definition, err = dictionary.Search("second")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(definition)

	err = dictionary.UpdateDefinition("first", "wordword")
	if err != nil {
		log.Fatal(err)
	}

	updatedWord, _ := dictionary.Search("first")
	fmt.Println(updatedWord)

	err = dictionary.DeleteWord("first")
	if err != nil {
		log.Fatal(err)
	}

}
