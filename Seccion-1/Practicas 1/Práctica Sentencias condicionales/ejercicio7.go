/* Realice un algoritmo para determinar si dos números son amigos. Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
   Sean X1, X2, XN todos divisores propios de X
   Sean Y1, Y2, YN todos divisores propios de Y
   Si X e Y son amigos entonces X1 + x2 +…+ XN es igual a Y e Y1 + Y2 +…+ XN es igual a X
   Ejemplo:
   El par (220, 284), ya que:
   Los divisores propios de 220 son 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 y 110, que suman 284.
   Los divisores propios de 284 son 1, 2, 4, 71 y 142, que suman 220. */

package main

import "fmt"

func obtenerDivisoresPropios(numero int) []int {
	divisores := []int{1} 
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			divisores = append(divisores, i)
			if i != numero/i {
				divisores = append(divisores, numero/i)
			}
		}
	}
	return divisores
}

func sumaDivisoresPropios(numero int) int {
	divisores := obtenerDivisoresPropios(numero)
	suma := 0
	for _, divisor := range divisores {
		suma += divisor
	}
	return suma
}

func sonAmigos(num1, num2 int) bool {
	return sumaDivisoresPropios(num1) == num2 && sumaDivisoresPropios(num2) == num1
}

func main() {
	
	num1, num2 := 221, 284

	if sonAmigos(num1, num2) {
		fmt.Printf("%d y %d son números amigos.\n", num1, num2)
	} else {
		fmt.Printf("%d y %d no son números amigos.\n", num1, num2)
	}
}
