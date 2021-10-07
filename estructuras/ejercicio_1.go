package main

import (
	"fmt"
)

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

// func imprimirClases(course Course)
func (course Course) ImprimirClases() {
	text := "Las clases son "
	for _, v := range course.Classes {
		text += v + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (course *Course) ChangePrice(p float64) {
	course.Price = p
}

func main() {
	matematica := Course{
		Name:    "Matematica desde Cero",
		Price:   45.50,
		IsFree:  false,
		UserIds: []uint{12, 56, 32},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Suma",
			3: "Resta",
		},
	}

	css := Course{
		Name:   "CSS",
		IsFree: true,
	}

	js := Course{}
	js.Name = "Javascrip"
	js.Price = 250.0
	js.Classes = map[uint]string{1: "Intro"}

	fmt.Println(matematica.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)
	// imprimirClases(matematica)
	matematica.ImprimirClases()
	matematica.ChangePrice(500.0)
	fmt.Println(matematica.Price)
}
