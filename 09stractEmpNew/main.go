package main

import (
	"bufio"
	"fmt"
	"os"
) 

	type Employee struct {
		name   string
		age    int
		job    string
		salary int
	}

	func main() {
		
		var pers1 Employee
		var pers2 Employee
		var pers3 Employee
		var pers4 Employee

	

	pers1.name = "Mursi"
	pers1.age = 25
	pers1.job = "Software Engineer"
	pers1.salary = 50000	

	pers2.name = "Kobi"
	pers2.age = 30
	pers2.job = "Data Scientist"
	pers2.salary = 60000			

	pers3.name = "David"
	pers3.age = 28
	pers3.job = "Product Manager"
	pers3.salary = 70000	


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter employee name: ")
	pers4.name, _ := reader.ReadString('\n')
	fmt.Print("Enter employee age: ")
	pers4.age, _ := reader.ReadString('\n')
	fmt.Print("Enter employee job: ")
	pers4.job, _ := reader.ReadString('\n')
	fmt.Print("Enter employee salary: ")




	PrintPerson(pers1)
	PrintPerson(pers2)
	PrintPerson(pers3)
	PrintPerson(pers4)



	}
	

	func PrintPerson(pers Employee) {
		fmt.Println("Name: ", pers.name)
		fmt.Println("Age: ", pers.age)
		fmt.Println("Job: ", pers.job)
		fmt.Println("Salary: ", pers.salary)
		fmt.Println("---------------------")
	}

	


