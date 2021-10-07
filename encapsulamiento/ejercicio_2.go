package main

import (
	"fmt"

	"github.com/Arthur98/poo/encapsulamiento/curso"
)

func main() {
	matematica := curso.New("Matematica", 0, true)
	matematica.SetClasses(map[uint]string{1: "Introduccion", 2: "Nivelacion"})
	matematica.Classes()
	fmt.Printf("%+v\n", matematica)
	fmt.Println(matematica.Classes())
	matematica.ImprimirClases()
}
