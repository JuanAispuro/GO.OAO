package main

import "fmt"

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T", i, i)

}

func main() {
	// var e exampler
	// e.x()
	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Juan")
}
