package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}

type Persona struct {
	name string
}

func NewPersona(name string) *Persona {
	return &Persona{name}
}

func (p *Persona) Get() string {
	return p.name
}

// Para actualizar la informacion es requerido que sea de tipo puntero.
func (p *Persona) Set(name string) {
	p.name = name
}

func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPersona("Arthur")
	Exec(p, "Charly")
}
