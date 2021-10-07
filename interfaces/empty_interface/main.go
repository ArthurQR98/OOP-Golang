package main

import (
	"fmt"
	"strings"
)

// Interfaces nulas
type exampler interface {
	x()
}

// Interfaces vacias - satisface cualquier tipo
func wrapper(i interface{}) {
	// fmt.Printf("Valor: %v , Tipo: %T\n", i, i)
	// validar el tipo (si es que lo quieres para un tipo especifico y hacer algo especifico con eso)
	// v, ok := i.(string)
	// if ok {
	// 	fmt.Printf("\t%s\n", strings.ToUpper(v))
	// }

	// TypeSwitch
	switch v := i.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("Valor: %v , Tipo: %T\n", i, i)
	}
}

func main() {
	var e exampler
	fmt.Printf("Valor: %v , Tipo: %T\n", e, e)
	// e.x()
	wrapper(12)
	wrapper(false)
	wrapper("Arthur")
	wrapper("Charly")
}
