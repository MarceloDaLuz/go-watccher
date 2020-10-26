package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		comandos()
	}
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
		monitorando()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo")
		saindoComSeguranca()
	default:
		fmt.Println("Comando não encontrado")
	}
}

func monitorando() {
	res, _ := http.Get("https://www.alura.com.br/")
	if res.StatusCode == 200 {
		fmt.Println("Alura está online!")
	} else {
		fmt.Println("Alura está offline!")
	}

}

func saindoComSeguranca() {
	os.Exit(0)
}
