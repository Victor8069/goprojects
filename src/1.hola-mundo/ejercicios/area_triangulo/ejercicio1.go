package main

import (
	"fmt"
	"math"
)

const precision = 2

func main() {
	var lado1, lado2 float64

	fmt.Println("Ingrese el lado 1: ")
	fmt.Scanln(&lado1)

	fmt.Println("Ingrese el lado 2: ")
	fmt.Scanln(&lado2)

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa
	s := perimetro / 2
	area := math.Sqrt(s * (s - lado1) * (s - lado2) * (s - hipotenusa))
	altura := (2 * area) / hipotenusa

	fmt.Printf("Hipotenusa del triangulo: %.*f \n", precision, hipotenusa)
	fmt.Printf("Perimetro del triangulo: %.*f \n", precision, perimetro)
	fmt.Printf("Area del triangulo: %.*f \n", precision, area)
	fmt.Printf("Altura del triangulo: %.*f \n", precision, altura)

}
