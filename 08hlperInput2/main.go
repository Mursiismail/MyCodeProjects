package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
) 

	func greet(name string) string {
		return "Hello, " + name + "!"
	}

	func main() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("what is your name? ")
		nameinput,_ := reader.ReadString('\n')
		nameinput = strings.TrimSpace(nameinput)
		massage := greet(nameinput)
		fmt.Println(massage)
	}

		