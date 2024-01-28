/*Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.*/

package main

import ("fmt")


func main() {

	

	var num int 

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println("El numero ingresado es igual a 0")
	} 	else if  num > 0 {
		fmt.Println("El numero ingresado es mayor a 0")
	} 	else {
		fmt.Println("el numero es menor a 0")
	}
}