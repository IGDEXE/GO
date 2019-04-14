// Go By Example
// Arrays
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Chama a funcao principal
func main() {
	// Cria um vetor com cinco posicoes, comecando por 0
	var a [5]int
	fmt.Println("Vazio:", a)
	// Podemos definir um valor indicando um indice
	a[4] = 100
	fmt.Println("Define:", a)
	fmt.Println("Pega:", a[4])
	// Usamos LEN para verificar o tamanho
	fmt.Println("Tamanho:", len(a))
	// Pode declarar todos os valores ao criar
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lotado:", b)
	// Eles sao unidimensionais, mas pode usar estruturas para torna-los bidimensionais
	var doisV [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			doisV[i][j] = i + j
		}
	}
	fmt.Println("Dois valores:", doisV)
}
