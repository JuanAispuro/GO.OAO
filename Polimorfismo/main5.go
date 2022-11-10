package main

import "fmt"

type Storager interface {
	Get() string //getter
	Set(string)  //setter
}
type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}
func (p *Person) Get() string {
	return p.name
}

// Va a actualizar el nombre de la persona
// Para que se actualice el mismo es con el puntero (*)
func (p *Person) Set(name string) {
	p.name = name
}
func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}
func main() {
	p := NewPerson("Juan")
	Exec(p, "Armando")
}
