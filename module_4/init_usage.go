package main

import "fmt"

var test = 1

func init() {
	test++
	fmt.Println("First init called.", test)
}

func init() {
	test++
	fmt.Println("Second init called.", test)
}
func main() {
	test++
	fmt.Println("Main function called.", test)
}
