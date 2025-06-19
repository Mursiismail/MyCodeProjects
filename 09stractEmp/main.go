package main
import (
	"fmt") 

	type Employee struct {
		Name   string
		age    int
		job    string
		salary int
	}

	func main() {
		
		var pers1 Employee
		var pers2 Employee
		var pers3 Employee

	

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

	PrintPerson(pers1)
	PrintPerson(pers2)
	PrintPerson(pers3)


	func PrintPerson(pers Employee) {
		fmt.Println("Name: ", pers.name)
		fmt.Println("Age: ", pers.age)
		fmt.Println("Job: ", pers.job)
		fmt.Println("Salary: ", pers.salary)
		fmt.Println("---------------------")
	}


