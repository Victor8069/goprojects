package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Victor"
	p.edad = 24
	p.correo = "victor@gmail.com"
	fmt.Println(p)

	p2 := Persona{"Cristiano", 40, "cristiano@gmail.com"}
	fmt.Println(p2)
}
