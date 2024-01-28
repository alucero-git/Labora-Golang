package main

import (
	"fmt"
)

func main() {
	
var planetName string
var numberOfMoons int
const satelite int = 100
var planen1, planet2, planet3 = "Mercurio", "Venus", "Tierra" 

planet4, planet5, planet6 := "jupiter", "Saturno", "Luna"


planetName = "Earth"
numberOfMoons = 1



fmt.Println(planetName, numberOfMoons, satelite, planet3, planen1, planet2, planet4, planet5, planet6)

fmt.Fprintln(planen1)
}