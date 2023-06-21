package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			showLogs()
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
	websites := getLinesFromFile()
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
		registerLog(website, true)
	} else {
		fmt.Println("Site:", website, "Falha")
		registerLog(website, false)
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

func getLinesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Falha ao abrir arquivo de sites")
		os.Exit(-1)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registerLog(website string, status bool) {
	logFile, err := os.OpenFile("events.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Falha ao abrir arquivo de logs")
	}

	if _, err := logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " - ONLINE: " + strconv.FormatBool(status) + "\n"); err != nil {
		logFile.Close()
	}

	if err := logFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func showLogs() {
	file, err := os.ReadFile("events.log")
	if err != nil {
		fmt.Println("Falha ao abrir arquivo de logs: ", err)
		os.Exit(-1)
	}

	fmt.Println(string(file))
}
