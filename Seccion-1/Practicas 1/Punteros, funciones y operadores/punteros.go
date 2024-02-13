// Escribir un programa en Go que declare dos variables enteras a y b con valores iniciales de 10 y 20, respectivamente. A continuación, declarar un puntero ptrA que apunte a la dirección de memoria de a. Utilizar el puntero ptrA para intercambiar los valores de a y b sin utilizar una variable auxiliar. Finalmente, imprimir los valores de a y b por pantalla para comprobar que se han intercambiado correctamente.
// Valores iniciales: a = 10 , b = 20
// Valores intercambiados: a = 20 , b = 10

package main

import "fmt"

func main() {

	a := 10
	b := 20

	ptrA := &a

	*ptrA, b = b, *ptrA

	fmt.Printf("Después del intercambio:\n")
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)

}
