package main

import "github.com/Samantha-Grace/udemygobeginner/first"

func main() {
	first.Greet()
	//private to package - cannot be accessed from outside package
	first.subGreet()
}
