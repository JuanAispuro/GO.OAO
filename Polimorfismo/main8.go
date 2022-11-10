package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {

	//esto es un type switch con type identifica que puede ser varios tipos.
	switch v := i.(type) {
	case string:
		fmt.Println("\t%s\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("Valor: %v, Tipo: %T", i, i)
	}
}

func main() {
	// var e exampler
	// e.x()
	wrapper(12)
	wrapper("juan")
	wrapper(14.32)
	wrapper(false)
	wrapper("Juan")
}
