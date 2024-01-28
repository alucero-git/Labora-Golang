/* Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.
Sea X el número.
X = S1 + S2 + S3 + S4 + S5
Donde
0 ≤ S1 ≤ 50
0 ≤ S2 ≤ 50
0 ≤ S3 ≤ 600
0 ≤ S4 ≤ 800
0 ≤ S5 < (Infinito)
De esta forma, si X vale 120, entonces
S1 = 50, S2 = 50, S3 = 20, S4 = 0 y S5 = 0
Si X vale 860, entonces
S1 = 50, S2 = 50, S3 = 600, S4 = 160 y S5 = 0 */

package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&x)

	S1, S2, S3, S4, S5 := descomponerNumero(x)
	fmt.Printf("Para X = %d, s1 = %d, s2 = %d, s3 = %d, s4 = %d, s5 = %d\n", x, S1, S2, S3, S4, S5)
}

func descomponerNumero(x int) (int, int, int, int, int) {
	limite1 := 50
	limite2 := 50
	limite3 := 600
	limite4 := 800

	s1, s2, s3, s4, s5 := 0, 0, 0, 0, 0

	if x > limite1 {
		s1 = limite1
		x -= limite1
	} else {
		s1 = x
		x = 0
	}

	if x > limite2 {
		s2 = limite2
		x -= limite2
	} else {
		s2 = x
		x = 0
	}

	if x > limite3 {
		s3 = limite3
		x -= limite3
	} else {
		s3 = x
		x = 0
	}

	if x > limite4 {
		s4 = limite4
		x -= limite4
	} else {
		s4 = x
		x = 0
	}

	s5 = x

	return s1, s2, s3, s4, s5
}
