package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Número:", num)

	result, err2 := divide(10, 2)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Resultado:", result)
}
