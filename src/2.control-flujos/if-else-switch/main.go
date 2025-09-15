package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Buenos días!")
	} else if t.Hour() < 19 {
		fmt.Println("¡Buenas tardes!")
	} else {
		fmt.Println("¡Buenas noches!")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> MacOS")
	default:
		fmt.Println("Go run -> Otro OS")

	}
}
