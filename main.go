package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/juniornelson123/encurtador-url/shortener"
)

func main() {
	clear()
	ui()
}

func ui() {
	scanner3 := bufio.NewScanner(os.Stdin)
	var action string

	for action != "2" {

		fmt.Println("*****************Encurtador de Url*******************")
		fmt.Println("1 - Encurtar Url")
		fmt.Println("2 - Sair")
		fmt.Println("Selecione uma ação: ")

		scanner3.Scan()
		action = scanner3.Text()
		switch {
		case action == "1":
			shortUrl()
			break
		case action == "2":
			fmt.Println("Saindo...")
			os.Exit(200)
			break

		default:
			main()
			fmt.Println("Valor invalido")
			break
		}

	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func shortUrl() {

	scanner := bufio.NewScanner(os.Stdin)
	// scanner1 := bufio.NewScanner(os.Stdin)

	var option string
	var url string

	for option != "q" {

		fmt.Println("Informa a url que deseja encurtar: ")

		scanner.Scan()
		url = scanner.Text()

		c, err := shortener.Short(url)
		errorFunc := <-err
		if errorFunc.Error() != "false" {

			fmt.Printf("\n\nErro: %s\n\n", errorFunc)
			shortUrl()
		} else {
			fmt.Printf("\n\nUrl Encurtada: %s\n\n", <-c)
			restart()
		}

	}

}

func restart() {
	scanner3 := bufio.NewScanner(os.Stdin)
	var action string

	fmt.Println("Deseja Realizar mais alguma ação ?")
	fmt.Println("1 - Menu Principal")
	fmt.Println("2 - Sair")

	scanner3.Scan()
	action = scanner3.Text()
	switch {
	case action == "1":
		main()
		break
	case action == "2":
		fmt.Println("Saindo...")
		os.Exit(200)
		break

	default:
		fmt.Println("Valor invalido")
		restart()
		break
	}

}
