package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Función para verificar si un número es capicúa
func esCapicua(numero int) bool {
	numeroComoCadena := strconv.Itoa(numero)
	return esPalindromo(numeroComoCadena)
}

// Función para verificar si una cadena es palíndromo
func esPalindromo(cadena string) bool {
	cadena = strings.ToLower(strings.ReplaceAll(cadena, " ", ""))

	for i, j := 0, len(cadena)-1; i < j; i, j = i+1, j-1 {
		if cadena[i] != cadena[j] {
			return false
		}
	}

	return true
}

func main() {
	// Ingresar un número o una cadena
	var entrada string
	fmt.Print("Ingrese un número o una cadena: ")
	fmt.Scan(&entrada)

	// Intentar convertir la entrada a un número entero
	numero, err := strconv.Atoi(entrada)
	if err == nil {
		// Verificar si el número es capicúa
		if esCapicua(numero) {
			fmt.Println("El número es capicúa.")
		} else {
			fmt.Println("El número no es capicúa.")
		}
	} else {
		// Verificar si la cadena es palíndromo
		if esPalindromo(entrada) {
			fmt.Println("La cadena es un palíndromo.")
		} else {
			fmt.Println("La cadena no es un palíndromo.")
		}
	}
}