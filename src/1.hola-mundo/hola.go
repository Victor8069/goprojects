package main

import (
	"fmt"
	// "strconv"

	"rsc.io/quote"
)

// Declaración de constantes
const pi float32 = 3.1415

const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	//Declaración de variables
	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName = "Victor"
	// 	lastName  = "Zapata"
	// 	age       = 24
	// )

	// firstName, lastName, age := "Victor", "Zapata", 24

	// fmt.Print(firstName, lastName, age)
	// fmt.Print(pi)
	// fmt.Print(x, y, z, w)
	// fmt.Print(Viernes)

	// fullName := "Alex Roe \t(alias \"roelcode\")\n"
	// fmt.Print(fullName)

	// var a byte = 'a'
	// fmt.Print(a)

	// s := "hola"
	// fmt.Print(s[3])

	// var integer16 int16 = 50
	// var integer32 int32 = 100

	// fmt.Print(int32(integer16) + integer32)

	// s := "100"
	// i, _ := strconv.Atoi(s)
	// fmt.Print(i + i)

	// n := 42
	// s = strconv.Itoa(n)
	// fmt.Print(s)
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	fmt.Println(greeting)
}
