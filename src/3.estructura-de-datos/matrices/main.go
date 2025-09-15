package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 10
	a[1] = 20
	fmt.Println(a)

	var b = [...]int{10, 20, 40, 80, 160}
	for i := 0; i < len(a); i++ {
		fmt.Println(b[i])
	}

	for index, value := range b {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}

	var matriz = [3][3]int{{2, 4, 8}, {16, 32, 64}, {128, 256, 512}}
	fmt.Println(matriz)

}
