package main

import "fmt"

type Persona struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Telefono int
}

func main() {

	persona1, persona2 := datosPersonas()

	primerImpresion(persona1, persona2)

	cambiarCiudad(&persona1, "Rosario")

	impresionPostCambioCiudad(persona1, persona2)
}

func impresionPostCambioCiudad(persona1 Persona, persona2 Persona) {
	fmt.Println("Después de cambiar la ciudad:")
	fmt.Printf("Persona 1: Nombre: %s, Edad: %d, Ciudad: %s, Teléfono: %d\n", persona1.Nombre, persona1.Edad, persona1.Ciudad, persona1.Telefono)
	fmt.Printf("Persona 2: Nombre: %s, Edad: %d, Ciudad: %s, Teléfono: %d\n", persona2.Nombre, persona2.Edad, persona2.Ciudad, persona2.Telefono)
}

func datosPersonas() (Persona, Persona) {
	persona1 := Persona{
		Nombre:   "Juan",
		Edad:     30,
		Ciudad:   "Buenos Aires",
		Telefono: 22334455,
	}

	persona2 := Persona{
		Nombre:   "Maria",
		Edad:     25,
		Ciudad:   "Madrid",
		Telefono: 11223344,
	}
	return persona1, persona2
}

func primerImpresion(persona1 Persona, persona2 Persona) {
	fmt.Println("Persona 1:")
	fmt.Printf("Nombre: %s, Edad: %d, Ciudad: %s, Teléfono: %d\n", persona1.Nombre, persona1.Edad, persona1.Ciudad, persona1.Telefono)

	fmt.Println("Persona 2:")
	fmt.Printf("Nombre: %s, Edad: %d, Ciudad: %s, Teléfono: %d\n", persona2.Nombre, persona2.Edad, persona2.Ciudad, persona2.Telefono)
}

func cambiarCiudad(p *Persona, nuevaCiudad string) {
	if p.Ciudad != nuevaCiudad {
		p.Ciudad = nuevaCiudad
		fmt.Printf("La ciudad de %s ha sido actualizada a %s.\n", p.Nombre, nuevaCiudad)
	} else {
		fmt.Printf("La ciudad de %s ya es %s. No se ha realizado ninguna actualización.\n", p.Nombre, nuevaCiudad)
	}
}
