package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// Un metodo pertenece a una estructura con la referencia *
func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println(x)
	fmt.Println(p)

	editar(&x)
	fmt.Println(x)

	pe := Persona{"Victor", 24, "victor@email.com"}
	pe.diHola()
}

func editar(x *int) {
	*x = 20
}
