package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main():")
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")
	// log.Fatal("¡Oye, soy un registro de errores!")
	// log.Panic("¡Oye, soy un registro de errores!")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("!Oye, soy un Log!")
}
