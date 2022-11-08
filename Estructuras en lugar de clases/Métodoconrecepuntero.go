package main
import "fmt"

func main(){
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
	(*Go).PrintClasess()
	// (&Go).ChangePrice(67.78)~
	// fmt.Println(Go.Price)
}