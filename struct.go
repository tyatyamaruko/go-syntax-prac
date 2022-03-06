package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func main() {
	// var p1 Person
	// p1.SetPerson("Yamada", 25)
	// name, age := p1.GetPerson()
	// fmt.Println(name, age)
	// fmt.Println(p1.name)
	// fmt.Println(p1.age)
	// p2 := Person{"Tanaka", 11}
	// fmt.Println(p2.GetPerson())

	p1 := P1{"山田太郎"}
	book := Book{title: "我輩は猫である"}

	PrintOut(p1)
	PrintOut(book)
}

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

type P1 struct {
	name string
}

func (p P1) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}
