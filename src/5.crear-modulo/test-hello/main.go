package main

import (
	"fmt"

	"github.com/VictorZ-IP/greetings"
)

func main() {
	message, err := greetings.Hello("Victor")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)
}
