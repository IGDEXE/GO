// Go By Example
// Switch
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Importa o pacote TIME
import "time"

// Chama a funcao principal
func main() {
	// Estrutura basica
	i := 2
	fmt.Print("Escreva ", i, " como ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")
	}
	// Pode usar virgulas para separar, e com DEFAULT, definir um valor padrao
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Final de semana")
	default:
		fmt.Println("Dia util")
	}
	// Sem uma expressao, funciona como IF-ELSE
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Antes do meio dia")
	default:
		fmt.Println("Depois do meio dia")
	}
	// Pode utilizar para comparar valores
	oQueSou := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Sou um tipo booleano")
		case int:
			fmt.Println("Sou um tipo inteiro")
		default:
			fmt.Printf("Sou um tipo desconhecido %T\n", t)
		}
	}
	oQueSou(true)
	oQueSou(1)
	oQueSou("Ola")
}
