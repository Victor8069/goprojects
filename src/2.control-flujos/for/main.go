package main

import "fmt"

func main() {
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			break
		}
		if j == 3 {
			continue
		}
		fmt.Println(j)
	}
}
