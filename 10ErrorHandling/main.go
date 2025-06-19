package main

import (
	"fmt"
	"os"
)
	
func main() { 


	result,err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	 fmt.Println("File content:", string(result))


}