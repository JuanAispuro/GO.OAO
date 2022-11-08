package main
/*
import "fmt"

//Type nombre struct {}
type Course struct {
	Name string
	Price float64
	IsFree bool 
	UsersID []uint
	Clasess map[uint]string
}
//Indicamos que es un metodo para que se vincule con el struct
//En Go no es buena practica utilizar los {this}
func (c Course) PrintClasess() {
	text := "Las clases son: "
	//forange
	for _, class := range c.Clasess{
		text += class + ", "
	}
	//Ponemos un slice para el comienzo del string hasta el final menos 2 para el espacio.
	fmt.Println(text[:len(text)-2])
}

func main(){
	//Tambien no es necesario poner el identificador aun lado simplemente siguiendo el orden de entrada.
	
		"Go Programación Orientada a Objetos",
		15.99,
		false,
		[]uint{12,56,89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
	
	//Variable Go
	Go := Course{
		Name: "Go Programación Orientada a Objetos",
		Price: 15.99,
		IsFree: false,
		UsersID: []uint{12,56,89},
		Clasess:map[uint]string{
			1: "Introducción" ,
			2: "Estructuras" ,
			3: "Maps" ,
		},
	}
	Go.PrintClasess()
	//Si quiero nomas poner algunos valores si queremos usar pocos, se tienen que inicialziar los nombres.
	/*
	css := Course{
		Name: "CSS",
		IsFree: true,
	}
	fmt.Printf("%+v\n", css)
	//Mas corto
	js := Course {}
	js.Name = "JavaScript"
	js.UsersID = []uint{12,67}
	fmt.Printf("%+v", js)
	
}
*/