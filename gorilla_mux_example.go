package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	if len(books) == 0 {
		_, _ = io.WriteString(w, "No books found")
		return
	}

	s, _ := json.Marshal(books)
	if _, err := io.WriteString(w, string(s)); err != nil {
		panic(err)
	}

}
func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	var book Book
	b, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b, &book); err != nil {
		panic(err)
	}
	books = append(books, book)
	fmt.Printf("%+v", book)
	if _, err := io.WriteString(w, "New book added"); err != nil {
		panic(err)
	}
}
func main() {
	r := mux.NewRouter()
	//GET Request
	r.HandleFunc("/books", getBooks).Methods(http.MethodGet)

	//POST Request
	r.HandleFunc("/books", addBook).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}
