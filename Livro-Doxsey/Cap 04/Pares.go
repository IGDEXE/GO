// Numeros pares

package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println("O numero ", i, " é par")
		} else {
			fmt.Println("O numero ", i, " é impar")
		}
		i++
	}
}
