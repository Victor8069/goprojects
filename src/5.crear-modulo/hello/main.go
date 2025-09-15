package main

import (
	"fmt"
	"log"

	"github.com/VictorZ-IP/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Victor", "Cristiano", "Adrian"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Victor")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(messages)
}
