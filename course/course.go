package course

import "fmt"

//Type nombre struct {}
type course struct {
	name    string
	price   float64
	Isfree  bool
	usersID []uint
	clasess map[uint]string
}

//Para usar getters no es necesari poner get solo que se llame igual con mayuscula.

//setters
func (c *course) SetClasses(clases map[uint]string) {
	c.clasess = clases
}

//Name
func (c *course) SetName(name string) { c.name = name }
func (c *course) Name() string        { return c.name }

//Price
func (c *course) SetPrice(price float64) { c.price = price }
func (c *course) Price() float64         { return c.price }

//Isfree
func (c *course) SetIsFree(isFree bool) { c.isFree = isFree }
func (c *course) isFree() bool          { return c.isFree }

//Setusers
func (c *course) SetUserIDs(userIDs []uint) { c.userIDs = userIDs }
func (c *course) UserIDs() []uint           { return c.userIDs }

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

//Mandamos una copia, pero pasamos un puntero para que mande el bueno.
//Si queremos actualizar una estructura con el (*)
/*
	func (c *course) changePrice(price float64) {
		c.price = price
	}
*/

/* Mejor poner todo de tipo puntero */

//tipo de dato que vamos a vincularlo.
//Si no queremos actualizar nada es con el normal.
func (c *course) PrintClasess() {
	text := "Las clases son: "
	for _, class := range c.clasess {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
