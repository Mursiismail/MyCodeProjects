package main

import "fmt"

func main() {

	var age int
	var name string
	var funfact string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter a fun fact about yourself: ")
	fmt.Scanln(&funfact)

	fmt.Println("Hello yor nqme is ", name)
	fmt.Println("You are ", age, " years old")
	fmt.Println("A fun fact about you is: ", funfact)	
}
