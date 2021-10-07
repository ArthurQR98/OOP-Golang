package main

import "fmt"

type course struct {
	name string
}

// declaracion de alias
type myAlias = course

// declaracion de tipo
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "Verdadero"
	}
	return "Falso"
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

func main() {
	c := newCourse{name: "Golang"}
	// c.Print() // newCourse no tiene esa funcion ya que otra tipo
	fmt.Printf("El tipo es : %T\n", c)

	var b newBool = true
	fmt.Println(b.String())
}
