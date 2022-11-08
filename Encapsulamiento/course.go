package main

import "fmt"

//Type nombre struct {}
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UsersID []uint
	Clasess map[uint]string
}

//Mandamos una copia, pero pasamos un puntero para que mande el bueno.
//Si queremos actualizar una estructura con el (*)
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

/* Mejor poner todo de tipo puntero */

//tipo de dato que vamos a vincularlo.
//Si no queremos actualizar nada es con el normal.
func (c *Course) PrintClasess() {
	text := "Las clases son: "
	for _, class := range c.Clasess {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
