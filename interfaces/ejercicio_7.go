package main

type Person struct {
	Name string
}

// Usando la interface Stringer y el metodo String para modificar la salida
func (p Person) String() string {
	return "Hola soy una persona , mi nombre es : " + p.Name
}

// func main() {
// 	p := Person{Name: "Arthur"}
// 	fmt.Println(p)
// }
