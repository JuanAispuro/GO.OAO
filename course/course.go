package course

import "fmt"

//Type nombre struct {}
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UsersID []uint
	Clasess map[uint]string
}

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

//Mandamos una copia, pero pasamos un puntero para que mande el bueno.
//Si queremos actualizar una estructura con el (*)
func (c *course) ChangePrice(price float64) {
	c.Price = price
}

/* Mejor poner todo de tipo puntero */

//tipo de dato que vamos a vincularlo.
//Si no queremos actualizar nada es con el normal.
func (c *course) PrintClasess() {
	text := "Las clases son: "
	for _, class := range c.Clasess {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
