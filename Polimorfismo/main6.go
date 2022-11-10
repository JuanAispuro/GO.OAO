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

func (p Cash) Pay() {
	fmt.Println("Pagado con Efectivo")
}
func (p CreditCard) Pay() {
	fmt.Println("Pagado con Tarjeta de Credito")
}

//Devolvemos el paymethod
func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		//retornamos nulo.
		return nil
	}
}

func main() {
	var method uint //variable
	//Solicitar
	fmt.Println("Digite un número para su metodo de pago.")
	fmt.Println("\t 1. Paypal 2: Efectivo 3: Tarjeta de Credito")
	_, err := fmt.Scanln(&method)
	//if para evitar que ponga un numero mal
	if err != nil {
		panic("debe digitar un número valido.")
	}
	if method > 3 {
		panic("debe digitar un numero de las opciones.")
	}

	//Llamamos a la fabrica
	paymethdo := Factory(method)
	paymethdo.Pay()
}
