package main

import (
	"fmt"
	"os"
	"text/template"
)

func checkError(e []error) {
	for _, err := range e {
		if err != nil {
			fmt.Printf("Error occurred: %v", err)
		}
	}
}

func main() {
	books := []Book{{Title: "Harry Potter", Author: "J.K. Rowling"}, {Title: "Wings of Fire", Author: "A P J Abdul Kalam"},
		{Title: "Harry Potter", Author: "J.K. Rowling"}}

	t := template.New("book desc")

	//parse some content and generate a template,
	//which is an internal representation
	t, _ = t.Parse("The book named {{.Title}} is written by {{.Author}}")

	//merge template ‘t’ with content of ‘p’
	err := t.Execute(os.Stdout, books[0])
	err2 := t.Execute(os.Stdout, books[1])

	template2 := template.New("bulk books")
	template2, _ = template2.Parse(`{{range .}}The book named {{.Title}} is written by {{.Author}}{{end}}`)
	// . beside range represents entire pipeline
	err3 := template2.Execute(os.Stdout, books)
	checkError([]error{err, err2, err3})
}
