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
			fmt.Println("Leaving program")
			os.Exit(0)
		default:
			fmt.Println("Don't know this command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "User"
	version := "1.0"
	fmt.Println("Hello,", name)
	fmt.Println("The program version is", version)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- View logs")
	fmt.Println("0- Exit program")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("The command chosen was", readedCommand)
	fmt.Println("")

	return readedCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	fmt.Println("")

	// websites := []string{"https://httpbin.org/status/404", "https://httpbin.org/status/200", "https://www.google.com.br"}

	websites := readArchiveWebsites()

	for i := 0; i < monitoring; i++ {
		for i, website := range websites {
			fmt.Println("Testing the website", i, ":", website)
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
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Response: ", website, "was loaded successfully!")
		registerLog(website, true)
	} else {
		fmt.Println("Response:", website, "is in trouble:", resp.StatusCode)
		registerLog(website, false)
	}
}

func readArchiveWebsites() []string {
	var websites []string

	archive, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
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
		fmt.Println("An error occurred:", err)
	}

	archive.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + website + " | online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func readLog() {
	archive, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Println(string(archive))
}
