package main

import (
	"fmt"
)

	type Student struct {
		Name  string
		school string
		year	int 
	}

	func main() {
		student1 := Student{
			Name:  "Mursi",
			school: "University of Science and Technology",
			year: 2024,
		}

		fmt.Println("tudent informaition")
		fmt.Println("Name:", student1.Name)
		fmt.Println("School:", student1.school)
		fmt.Println("Year:", student1.year)
	}
	

		// Accessing the fields directly
