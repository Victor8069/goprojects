package main

import (
	// "fmt"
	"library/animal"
)

func main() {
	// var myBook = book.Book{
	// 	Title:  "Moby Dick",
	// 	Author: "Herman Melville",
	// 	Pages:  300,
	// }

	// mathBook := book.NewBook("Maths", "Aristoteles", 500)

	// myBook.PrintInfo()
	// mathBook.PrintInfo()

	// mathBook.SetTitle("Pro Maths")
	// fmt.Println(mathBook.GetTitle())

	// myTextBook := book.NewTextBook("Comunicación", "Jaime Gamarra", 262, "Santillana SAC", "Secundaria")
	// myTextBook.PrintInfo()

	// book.Print(mathBook)
	// book.Print(myTextBook)

	// miPerro := animal.Perro{
	// 	Nombre: "Firulais",
	// }
	// miGato := animal.Gato{
	// 	Nombre: "Michi",
	// }

	// animal.HacerSonido(&miPerro)
	// animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Firulais"},
		&animal.Gato{Nombre: "Michi"},
		&animal.Perro{Nombre: "Rex"},
		&animal.Gato{Nombre: "Garfield"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
