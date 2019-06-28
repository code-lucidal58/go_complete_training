package main

import (
	"fmt"
	"os"
	"text/template"
)

type Book struct {
	Title  string
	Author string
}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	books := []Book{
		{Title: "Harry Potter", Author: "J.K. Rowling"},
		{Title: "Great Expectations", Author: "Charles Dickens"},
		{Title: "The Adventures of Huckleberry Finn", Author: "Mark Twain"},
	}

	// template1 for parsing data for only one struct
	template1 := template.New("book_description")

	//parse content to generate a template
	template1, _ = template1.Parse("The book named {{.Title}} is written by {{.Author}}\n")

	//merge template ‘template1’ with content of ‘books’
	err := template1.Execute(os.Stdout, books[0])
	CheckError(err)
	err = template1.Execute(os.Stdout, books[1])
	CheckError(err)

	//Pipeline of data
	t := template.New("template test")
	t = template.Must(t.Parse("\nThis is just static text. \n{{\"This is pipeline data - because it is evaluated " +
		"within the double braces.\"}} {{`So is this, but within reverse quotes.`}}\n"))
	_ = t.Execute(os.Stdout, nil)

	// template2 for parsing data for multiple structs in one go
	template2 := template.New("books_description")
	template.Must(template2.Parse(`{{range .}}The book named {{.Title}} is written by {{.Author}}
{{end}}`))

	// . beside range represents entire pipeline
	err = template2.Execute(os.Stdout, books)
	CheckError(err)

	// parsing template from a file
	template3 := template.New("books_file")
	template3, err = template3.ParseFiles("module_1/templates/text_template.gotmpl")
	CheckError(err)
	err = template3.Execute(os.Stdout, books)
	CheckError(err)
}
