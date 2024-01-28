/*Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.*/

package main

import (
	"fmt"
)

func suma(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func prom(num1, num2, num3 int) int {
	return (num1 + num2 + num3) / 3
}


func main() {

	var num1 int
	var num2 int
	var num3 int



	fmt.Print("Ingrese el primer valor de la temperatura: ")
	fmt.Scan(&num1)

	fmt.Print("Ingrese el segundo valor de la temperatura: ")
	fmt.Scan(&num2)

	fmt.Print("Ingrese el tercer valor de la temperatura: ")
	fmt.Scan(&num3)

	sumaTotal := suma(num1,num2,num3)
	promTotal := prom(num1,num2,num3)

	fmt.Printf("el valor de la suma es: %d\n", sumaTotal)
	fmt.Printf("el valor del promedio es: %d\n", promTotal)

}

