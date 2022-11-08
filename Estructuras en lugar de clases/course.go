package main

import "fmt"
type Course struct{
	Name string
	Price float64
	IsFree bool 
	UsersID []uint
	Clasess map[uint]string
}
//puntero
func (c *Course) ChangePrice (price float64){
	c.Price = price
}


func (c Course) PrintClasess() {
	text := "Las clases son: "
	//forange
	for _, class := range c.Clasess{
		text += class + ", "
	}
	//Ponemos un slice para el comienzo del string hasta el final menos 2 para el espacio.
	fmt.Println(text[:len(text)-2])
}
