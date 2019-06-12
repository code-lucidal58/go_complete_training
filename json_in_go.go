package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	// Single struct object
	book := Book{Title: "Harry Potter", Author: "J.K. Rowling"}
	b, _ := json.Marshal(book)
	err := ioutil.WriteFile("demo.json", b, 0644)
	checkError(err)

	fileData, err := ioutil.ReadFile("demo.json")
	var bookParsed Book
	err = json.Unmarshal(fileData, &bookParsed)
	checkError(err)
	fmt.Printf("%+v", bookParsed)

	// Array of struct objects
	books := []Book{{Title: "Harry Potter", Author: "J.K. Rowling"}, {Title: "Wings of Fire", Author: "A P J Abdul Kalam"},
		{Title: "Harry Potter", Author: "J.K. Rowling"}}
	b, _ = json.Marshal(books)
	err = ioutil.WriteFile("demo_array.json", b, 0644)
	checkError(err)

	fileData, err = ioutil.ReadFile("demo_array.json")
	var bookParsedArray []Book
	err = json.Unmarshal(fileData, &bookParsedArray)
	checkError(err)
	fmt.Printf("%+v", bookParsedArray)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
