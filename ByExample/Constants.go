// Go By Example
// Constants
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Importa o pacote MATH
import "math"

// Utilizamos 'const' para declarar constantes
const s string = "constante"

// Chama a funcao principal
func main() {
	// Escreve na tela o valor que salvamos na constante
	fmt.Println(s)
	// const funciona de forma muito parecida com var
	const n = 500000000
	// Eh possivel fazer operacoes matematicas
	const d = 3e20 / n
	fmt.Println(d)
	// Podemos definir um tipo
	fmt.Println(int64(d))
	// E utilizar operacoes de pacotes
	fmt.Println(math.Sin(n))
}
