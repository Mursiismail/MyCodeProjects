package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter you fav authar")
	favAut ,_  := reader.ReadString('\n')
	fmt.Println("your Favourite author is ", favAut)
	fmt.Print("Enter your favourite book Tittle")
	favBook, _ := reader.ReadString('\n')
	fmt.Println("Your favourite number is ", favBook)
	fmt.Print("Enter year of publication")
	yearOfPaplcation, _ := reader.ReadString('\n')
	fmt.Println("A fun fact about you is: ", yearOfPaplcation)
}	