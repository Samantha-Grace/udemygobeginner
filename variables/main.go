package main

import "fmt"

//package scope
var year = 2020

//block scope - must use if you declare
func main() {

	count := 10
	country := "India"
	var print = true

	fmt.Println(count)
	fmt.Println(country)
	fmt.Println(print)
}

/*
multi variable declaration
func main(

	count int
	country string
	print bool
	firstName string
)

*/
