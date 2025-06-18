package main

import (
	"fmt"
)



func greet(name string) string {
		return "Hello, " + name + "!"}

func main() { 
	massage := greet("Mursi")
	fmt.Println (massage)
} 

