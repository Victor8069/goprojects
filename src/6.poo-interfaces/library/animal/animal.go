package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func NewDog(name string) *Perro {
	return &Perro{
		Nombre: name,
	}
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

type Gato struct {
	Nombre string
}

func NewCat(name string) *Gato {
	return &Gato{
		Nombre: name,
	}
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau miau")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
