package curso

import "fmt"

// Variables exportados siempre tiene que empezar con Mayuscula y minuscula
// sino deseas acceder a el.
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	classes map[uint]string
}

// Como si fuera un constructor
func New(name string, price float64, isFree bool) *course {
	// para agregar un valor por default
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) Classes() map[uint]string {
	return c.classes
}

// func imprimirClases(course Course)
func (course *course) ImprimirClases() {
	text := "Las clases son "
	for _, v := range course.classes {
		text += v + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (course *course) ChangePrice(p float64) {
	course.Price = p
}
