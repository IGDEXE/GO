// Tipos e conversoes 
// Ivo Dias

// Pacote principal
package main

// Importa a biblioteca
import( 
	"fmt"
	"strconv"
)

// Funcao principal
func main() {
	idade := 22 // Declara uma variavel para a idade
	numero := "12" // Declara uma string
	idade_ToString := strconv.Itoa(idade) // Converte o valor inteiro para string
	numero_ToInt,_ := strconv.Atoi(numero) // Converte de string para int
	fmt.Println("Minha idade eh "+idade_ToString) // Exibe na tela
	fmt.Println(idade + numero_ToInt) // Exibe a soma com o numero convertido
}