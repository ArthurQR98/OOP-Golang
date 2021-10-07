package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

// Embebiendo interfaces (solo interfaces) composicion de interfaces
type GreeterByer interface {
	Greeter
	Byer
}

type People struct {
	Name string
}

func (p People) Greet() {
	fmt.Println("Hola soy", p.Name)
}

func (p People) Bye() {
	fmt.Println("Adios soy", p.Name)
}

type Texto string

func (t Texto) Greet() {
	fmt.Println("Hola soy", t)
}

func (t Texto) Bye() {
	fmt.Println("Adios soy", t)
}

func All(gb ...GreeterByer) {
	for _, v := range gb {
		v.Greet()
		v.Bye()
	}
}

// func GreeterAll(gs ...Greeter) {
// 	for _, g := range gs {
// 		g.Greet()
// 		fmt.Printf("\t Valor:%v, Tipo: %T\n", g, g)
// 	}
// }

// func BayAll(by ...Byer) {
// 	for _, b := range by {
// 		b.Bye()
// 		// fmt.Printf("\t Valor:%v, Tipo: %T\n", b, b)
// 	}
// }

// func main() {
// 	// var g Greeter = People{Name: "Arthur"}
// 	// var g Greeter = Texto("Arthur")
// 	// g.Greet()

// 	// var t Texto = "Deysi"
// 	// t.Greet()

// 	p := People{Name: "Arthur"}
// 	var t Texto = "Deysi"
// 	// GreeterAll(p, t)
// 	// BayAll(p, t)
// 	All(p, t)
// }
