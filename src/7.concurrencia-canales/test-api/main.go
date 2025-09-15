package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://api.github.com",
		"https://www.googleapis.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://api.somewhereintheinternet.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: %s está caído\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s está en funcionamiento\n", api)
}
