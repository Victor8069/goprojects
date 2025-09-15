package main

import "fmt"

// Función Variádica
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	return total
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	suma("Victor", 12, 45, 7852, 80, 80)
	imprimirDatos("Hola", 18, true, 3.14)
}
