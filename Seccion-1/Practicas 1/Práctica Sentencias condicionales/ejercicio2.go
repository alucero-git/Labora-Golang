/*Escribir un algoritmo que determine si un número es par.*/

package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Print("ingrese un numero: ")
	fmt.Scan(&num)

	if  num % 2 == 0 {
		fmt.Println("el numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

}
