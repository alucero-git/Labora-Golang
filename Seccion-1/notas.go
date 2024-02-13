/* Hagamos una aplicación Golang sencilla en la cual:

Creemos un arreglo llamado “notas” que tenga 5 números enteros que representen calificaciones del alumno
Declaremos una variable llamada “notas” de tipo arreglo de enteros
Construyamos un arreglo de 5 elementos y guardemos su referencia en la variable de ese tipo.
Pongamos 5 valores que representen calificaciones
Utilicemos los indices
Sumemos las 5 notas
Obtengamos el promedio */

package main

import "fmt"

func main() {
	var notas [5]float32 //= [5]int{4, 6, 8, 10, 6}

	notas[0] = 4
	notas[1] = 6
	notas[2] = 8
	notas[3] = 10
	notas[4] = 6

	var acumulador float32
	for i := 0; i < len(notas); i++ {

		acumulador += notas[i]

		fmt.Println(acumulador)

	}

	var promedio float32
	promedio = acumulador / float32(len(notas))

	fmt.Println("el promedio es: ", promedio)
}
