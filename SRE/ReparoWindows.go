// Reparar Windows com Go
// Ivo Dias

// Pacote principal
package main

// Importa a biblioteca
import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
	"path/filepath"
)

// Funcao principal
func main() {
	// Cria o LOG
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) // Verifica a pasta atual
	dir += "\\ReparoWindows.txt"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	exec.Command("cmd","/C","echo","Iniciando reparo do sistema",">> ",dir).Run()
	exec.Command("cmd","/C","echo","date", "/t",">> ",dir).Run()
	exec.Command("cmd","/C","echo","time /t",">> ",dir).Run()
	// Faz o reparo em si
	fmt.Println("O reparo do sistema pode demorar, favor nao fechar a tela ate ele completar") // Avisa sobre a demora do procedimento
	fmt.Println("Reparando o Windows") // Escreve na tela
	//exec.Command("cmd","/C","sfc /scannow").Run() // Usa o SFC para corrigir erros no Windows
	fmt.Println("Reparando a imagem do sistema") // Escreve na tela
	//exec.Command("cmd","/C","dism /online /cleanup-image /restorehealth").Run() // Usa o DISM para reparar a imagem do Windows
	fmt.Println("Reparo concluido") // Escreve na tela
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Aguarda uma entrada de dados para fechar a tela
}