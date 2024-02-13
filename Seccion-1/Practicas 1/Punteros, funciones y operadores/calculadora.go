/*
Escribir un programa que pida al usuario que ingrese dos números enteros. Luego, se debe crear una función llamada calcular que tome dos punteros a enteros como argumentos y calcule la suma, resta, multiplicación y división de los números apuntados por los punteros. Finalmente, se deben imprimir los resultados de las operaciones por pantalla.
la salida del programa si el usuario ingresa los números 10 y 5:
Ingrese el primer número: 10
Ingrese el segundo número: 5
Suma: 15
Resta: 5
Multiplicación: 50
División: 2.00
*/

package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&num1)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&num2)

	suma, resta, mult, div := calcular(&num1, &num2)

	resultados(suma, resta, mult, num2, div)

}

func resultados(suma int, resta int, mult int, num2 int, div float64) {
	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Resta: %d\n", resta)
	fmt.Printf("Multiplicación: %d\n", mult)
	if num2 == 0 {
		fmt.Println("División: No se puede dividir por cero")
	} else {
		fmt.Printf("División: %.2f\n", div)
	}
}

func calcular(a, b *int) (int, int, int, float64) {

	suma := *a + *b
	resta := *a - *b
	multiplicacion := *a * *b

	if *b == 0 {
		return suma, resta, multiplicacion, 0.0
	}
	division := float64(*a) / float64(*b)
	return suma, resta, multiplicacion, division

}
