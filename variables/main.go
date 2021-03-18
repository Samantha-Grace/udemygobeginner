package main

import "fmt"

//package scope
var name string

//block scope - must use if you declare
func main() {
	var count int
	var country string
	var print bool
	var firstName string

	var miles, age int

	//blank identifier because firstName is in the block but not used
	_ = firstName

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
