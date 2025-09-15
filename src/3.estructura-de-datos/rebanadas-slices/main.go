package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)

	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(diasSemana)

	// Crear rebanada a partir de otra
	diaRebanada := diasSemana[0:5]
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro dia")
	fmt.Println(diaRebanada)

	// Eliminar elementos
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	// Declarar slice
	nombres := make([]string, 5, 10)
	nombres[0] = "Alex"
	fmt.Println(nombres)

	// Copiar elementos de una rebanada a otra
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	copy(rebanada2, rebanada1)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)

}
