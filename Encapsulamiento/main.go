package main

func main() {
	//Tambien no es necesario poner el identificador aun lado simplemente siguiendo el orden de entrada.
	Go := &course.Course{
		"Go Programación Orientada a Objetos",
		15.99,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}
	//(*) Desreferenciacion
	Go.PrintClasess()
	//& obtener la memoria y direccion
	// Go.ChangePrice(67.12)
	// fmt.Println(Go.Price)
}
