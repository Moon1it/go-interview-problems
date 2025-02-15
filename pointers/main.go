package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

// person is a pointer to a Person struct
// when we try to change the value of the pointer itself,
// we are not changing the original pointer!

func changePerson(person *Person) {
	person = &Person{
		name: "John",
		age:  20,
	}
}

func main() {
	person := &Person{
		name: "Mike",
		age:  25,
	}

	fmt.Println(person.name)
	changePerson(person)
	fmt.Println(person.name)
}
