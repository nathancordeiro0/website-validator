package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func showIntroduction() {
	name := "Nathan"
	version := "23.2"
	fmt.Println("Olá,", name)
	fmt.Println("A versão desse programa é a", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("O comando escolhido foi", readedCommand)
	fmt.Println("")

	return readedCommand
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	// websites := []string{"https://httpbin.org/status/404", "https://httpbin.org/status/200", "https://www.google.com.br"}

	websites := readArchiveWebsites()

	for i := 0; i < monitoring; i++ {
		for i, website := range websites {
			fmt.Println("Testando o site", i, ":", website)
			testWebsite(website)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testWebsite(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Resposta: ", website, "foi carregado com sucesso!")
		registerLog(website, true)
	} else {
		fmt.Println("Resposta:", website, "está com problemas:", resp.StatusCode)
		registerLog(website, false)
	}
}

func readArchiveWebsites() []string {
	var websites []string

	archive, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(archive)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	archive.Close()

	return websites
}

func registerLog(website string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + website + " | online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func readLog() {
	archive, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(archive))
}

func main() {

	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			readLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}
