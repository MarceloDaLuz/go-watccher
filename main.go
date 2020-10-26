package main

import (
	"fmt"
	"os"
)

func main() {
	comandos()
}

func comandos() {
	fmt.Println("========== WATCCHER ==========")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
	fmt.Println("==============================")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramentos")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo")
		saindoComSeguranca()
	default:
		fmt.Println("Comando não encontrado")
	}
}

func saindoComSeguranca() {
	os.Exit(0)
}
