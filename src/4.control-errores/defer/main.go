package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Alex Roel"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
