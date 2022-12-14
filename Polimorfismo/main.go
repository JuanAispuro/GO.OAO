package main

import "fmt"

//Declaración de interface
type Greeter interface {
	Greet()
}

//estructura
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s", t)
}

func main() {
	var g Greeter = Person{Name: "Juan"}
	g.Greet()
	//Es un tipo texto no un string por lo que lo cambiaremos como:
	var t Greeter = Text("Texto")
	t.Greet()
}
