package mydict

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var errNotFound = errors.New("NOT FOUND")
var errWordExists = errors.New("WORD ALREADY EXISTS")

// types can have methods

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add word to dictionary
func (d Dictionary) AddWord(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word's definition
func (d Dictionary) UpdateDefinition(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		d[word] = def
	}
	return nil
}

// Delete a word
func (d Dictionary) DeleteWord(word string) error {
	defer fmt.Println("word", word, "deleted")

	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		delete(d, word)
	}
	return nil
}
