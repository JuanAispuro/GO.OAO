package main

import (
	"fmt"
)

// Declaración de interface
type Greeter interface {
	Greet()
}

// Metodo bye
type Byer interface {
	Bye()
}

// Interfaz completa
type GreeterByer interface {
	Greeter
	Byer
}

// estructura
type Person struct {
	Name string
}

func (p Person) String() string {
	//Con esto persona utilizo la interfaz stringer de format.
	return "Hola soy una persona y mi nombre es: " + p.Name
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s", p.Name)
}

type Text string

/*
	func GreetAll(gs ...Greeter) {
		for _, g := range gs {
			g.Greet() //valor tipo %v
			fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
		}
	}

	func ByeAll(bs ...Byer) {
		for _, b := range bs {
			b.Bye() //valor tipo %v
			fmt.Printf("\t Valor: %v, Tipo: %T\n", b, b)
		}
	}
*/
func (t Text) Greet() {
	fmt.Printf("Hola soy %s", t)
}

func (t Text) Bye() {
	fmt.Printf("Adios soy %s", t)
}
func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}
func main() {
	p := Person{Name: "Juan"}

	//Es un tipo texto no un string por lo que lo cambiaremos como:
	// var t Text = Text("Texto")
	// All(p, t)
	fmt.Println(p)
}
