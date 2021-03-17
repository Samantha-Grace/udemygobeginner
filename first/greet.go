package first

import "fmt"

// if greet is lowercase, it is private to this package
// Greet this is my first library package
func Greet() {
	fmt.Println("hello")
	//add sub here
	subGreet()

}

func subGreet() {
	fmt.Println("insidesub")

}
