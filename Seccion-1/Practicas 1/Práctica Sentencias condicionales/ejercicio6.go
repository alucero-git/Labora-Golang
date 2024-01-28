/*  Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
    1. 1030 segundos son 17 minutos con 10 segundos
    2. 12045 segundos son 3 horas, 20 minutos con 45 segundos
    3. 176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos. */

package main

import (
	"fmt"
	"time"
)

func main() {

	var segundos int

	fmt.Print("Ingrese los segundos que quiere calcular: ")
	fmt.Scan(&segundos)

	fmt.Printf("Para los segundos ingresados = %d son: dias: %d, horas: %d, minutos: %d, segundos: %d\n", segundos, obtenerDias(segundos), obtenerHoras(segundos), obtenerMinutos(segundos), obtenerSegundosRestantes(segundos))

}

func obtenerDias(segundos int) int {
	return int(time.Second * time.Duration(segundos) / (24 * time.Hour))
}

func obtenerHoras(segundos int) int {
	return int((time.Second * time.Duration(segundos) % (24 * time.Hour)) / time.Hour)
}

func obtenerMinutos(segundos int) int {
	return int((time.Second * time.Duration(segundos) % time.Hour) / time.Minute)
}

func obtenerSegundosRestantes(segundos int) int {
	return int(time.Second * time.Duration(segundos) % time.Minute / time.Second)
}
