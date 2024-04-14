package abstract_factory

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	Name string
	Age  int
}

func (p person) Greet() {
	fmt.Println("hello world")
}

func NewPerson(name string, age int) Person {
	return &person{
		Name: name,
		Age:  age,
	}
}
