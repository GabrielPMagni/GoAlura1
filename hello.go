package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay time.Duration = 2 * time.Minute

func main() {
	showProgramIntro()
	for {
		displayMenu()
		command := getMenuOption()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exbindo logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func startMonitoring() {
	fmt.Println("Monitorando")
	websites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}
	for {
		for _, website := range websites {
			testWebsite(website)
		}
		time.Sleep(delay)
		fmt.Println("")
	}
}

func testWebsite(website string) {
	response, err := http.Get(website)
	if response.StatusCode == 200 && err == nil {
		fmt.Println("Site:", website, "Sucesso")
	} else {
		fmt.Println("Site:", website, "Falha")
	}
}

func showProgramIntro() {
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func displayMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Print("")
}

func getMenuOption() int {
	var command int
	fmt.Scan(&command)
	return command
}
