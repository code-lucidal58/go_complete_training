package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	_, err := io.WriteString(w, "Hello world!")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
