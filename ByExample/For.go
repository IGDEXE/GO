// Go By Example
// For
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Chama a funcao principal
func main() {
	// Com uma condicao simples
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	// A classic initial/condition/after `for` loop.
	// Utilizando a estrutura Inicio / Condicao / Incremente
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// Sem uma condicao, eh preciso de um BREAK para parar o loop
	for {
		fmt.Println("loop")
		break
	}
	// Utilize CONTINUE para ir para o proximo ciclo
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
