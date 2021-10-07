package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}
type Cash struct{}
type CreditCard struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado por Paypal")
}

func (c Cash) Pay() {
	fmt.Println("Pagado por Cash")
}

func (cc CreditCard) Pay() {
	fmt.Println("Pagado con Tarjeta de credito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite un numero de metodo de pago")
	fmt.Println("\t (1) Paypal,(2) Cash,(3) Tarjeta de Credito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe digitar un metodo de pago valido")
	}

	if method > 3 {
		panic("debe digitar un metodo de pago valido")
	}

	methodPay := Factory(method)
	methodPay.Pay()
}
