package main

import (
	"fmt"
	"github.com/JuanAispuro/GO.OAO/tree/main/course"
)

func main() {
	Go := course.New("Go desde cero", 0, false)
	Go.UsersID = []uint{12, 56, 89}
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	//(*) Desreferenciacion
	Go.PrintClasess()
	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	//& obtener la memoria y direccion
	Go.ChangePrice(67.12)
	// fmt.Println(Go.Price)
}
