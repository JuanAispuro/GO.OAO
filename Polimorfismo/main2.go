package main

import (
	"fmt"
)

// Declaración de interface
type Greeter interface {
	Greet()
}

// estructura
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

type Text string

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet() //valor tipo %v
		fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
	}
}
func (t Text) Greet() {
	fmt.Printf("Hola soy %s", t)
}

func main() {
	p := Person{Name: "Juan"}

	//Es un tipo texto no un string por lo que lo cambiaremos como:
	var g Text = Text("Texto")
	GreetAll(p, g)
}
