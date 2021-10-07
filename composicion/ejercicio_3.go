package main

import (
	"fmt"

	"github.com/Arthur98/poo/composicion/customer"
	"github.com/Arthur98/poo/composicion/invoice"
	"github.com/Arthur98/poo/composicion/invoiceitem"
)

func main() {
	i := invoice.New("Peru", "Trujillo", customer.New("Arthur", "San Agustin", "938898701"), []invoiceitem.Item{invoiceitem.New(1, "Curso de Golang", 12.35), invoiceitem.New(1, "Curso de POO Golang", 50.55)})

	ii := invoice.New("Peru", "Trujillo", customer.New("Arthur", "San Agustin", "938898701"), []invoiceitem.Item{invoiceitem.New(1, "Curso de Typescript", 80.00), invoiceitem.New(1, "Curso de Amazon Web Servicies", 180.00)})

	iii := invoice.New("Peru", "Trujillo", customer.New("Arthur", "San Agustin", "938898701"), invoiceitem.NewItems(invoiceitem.New(1, "Curso de Typescript", 80.00), invoiceitem.New(1, "Curso de Amazon Web Servicies", 180.00)))

	i.SetTotal()
	ii.SetTotal()
	fmt.Printf("%+v\n", i)
	fmt.Printf("%+v\n", ii)
	fmt.Printf("%+v\n", iii)
}
