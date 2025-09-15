package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	// Agregar elemento al mapa
	colors["negro"] = "#000000"
	fmt.Println(colors)

	// Recuperar valor y comprobar si existe la clave
	if valor, ok := colors["rojo"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe esta clave")
	}

	delete(colors, "verde")
	fmt.Println(colors)

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}
