package main

import (
	"github.com/JuanAispuro/GO.OAO/tree/main/Encapsulamiento/course"
)

func main() {
	Go := course.New("Go desde cero", 0, false)
	Go.UsersID = []uint{12, 56, 89}
	Go.Clasess = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	},

		//(*) Desreferenciacion
		Go.PrintClasess()
	//& obtener la memoria y direccion
	Go.ChangePrice(67.12)
	// fmt.Println(Go.Price)
}
