/* Descripción del problema
Escribe un programa que calcule el área y el perímetro de un rectángulo. El programa debe pedir al usuario que introduzca la longitud y la anchura del rectángulo y luego calcular y mostrar el área y el perímetro.
Especificaciones:
Solicita al usuario que ingrese la longitud y la anchura del rectángulo.
Calcula el área del rectángulo (longitud * anchura).
Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
Muestra los resultados (área y perímetro) al usuario.
Ejemplo de Salida:
Ingrese la longitud del rectángulo: 5
Ingrese la anchura del rectángulo: 3
El área del rectángulo es: 15
El perímetro del rectángulo es: 16 */

package main

import (
	"fmt"
)

func main() {
	var lon int
	var anc int

	fmt.Print("Ingrese la longitud: ")
	fmt.Scan(&lon)

	fmt.Print("Ingrese la anchura: ")
	fmt.Scan(&anc)

areaTotal := area(lon,anc)
perimetroTotal := perimetro(lon,anc)

	fmt.Printf("El area es: %d y el perimetro es: %d", areaTotal, perimetroTotal)
}

func area(lon,anc int) int   {
	return (lon * anc)
}
func perimetro(lon, anc int) int  {
	return 2 * (lon + anc)
}