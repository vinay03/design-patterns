package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s. I am %d years old \n", p.name, p.age)
}
func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, i am too tired\n", p.name, p.age)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			name, age,
		}
	} else {
		return &person{
			name, age,
		}
	}
}

func main() {
	p := NewPerson("John", 33)
	p.SayHello()
}
