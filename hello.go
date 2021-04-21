package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibiIntroducao()

	for {
		exibiMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não Conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibiIntroducao() {
	nome := "Diego"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibiMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Digitei", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "esta OK")
	} else {
		fmt.Println("O", site, "retornou", resp.StatusCode, ":(")
	}
}
