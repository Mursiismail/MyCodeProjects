package main

import (
	"fmt"
)

	type Student struct {   // Defining a struct  ( type - Name - stract ) 
		Name  string
		school string
		year	int 
	}

	func main() {
		student1 := Student{
			Name:  "Mursi",
			school: "University of Science and Technology",
			year:  2024,}

			student2 := Student{
				Name:  "Kobi",
				school: "University of Science and Technology",
				year:  2024,}

				student3 := Student{
					Name:  "David",
					school: "University of Science and Technology",
					year:  2024,}

			student4 := Student{
				Name:  "valad",
				school: "University of Science and Technology",
				year:  2024,}

		

		fmt.Println( ('\n') , "student informaition")
		fmt.Println("Name:", student1.Name)
		fmt.Println("School:", student1.school)
		fmt.Println("Year:", student1.year)

		fmt.Println( ('\n') ,"student informaition")
		fmt.Println("Name:", student2.Name)
		fmt.Println("School:", student2.school)
		fmt.Println("Year:", student2.year)

		fmt.Println("student informaition")
		fmt.Println("Name:", student3.Name)
		fmt.Println("School:", student3.school)
		fmt.Println("Year:", student3.year)

		fmt.Println("student informaition")
		fmt.Println("Name:", student4.Name)
		fmt.Println("School:", student4.school)
		fmt.Println("Year:", student4.year)
	}
	

		// Accessing the fields directly
