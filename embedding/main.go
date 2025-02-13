package main

import "fmt"

type Person struct {
	name string
}

func (this *Person) Introduce() {
	fmt.Printf("I am a Person, my name is %s!\n", this.name)
}

type Employee struct {
	Person
	role string
}

func (this *Employee) Introduce() {
	fmt.Printf("I am a Employee, my name is %s! I'm working as %s\n", this.name, this.role)
}

func main() {
	e := Employee{
		Person: Person{
			name: "Vasily",
		},
		role: "Programmer",
	}
	e.Introduce()
	e.Person.Introduce()
}
