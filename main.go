package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	welcome()
	login, user := auth()
	if login && user != "" {

	}
	for {

	}
}

func auth() (bool, string) {
	var tentativas int = 0
	var usr string
	var pass string
	for {
		fmt.Println("Usuário:")
		fmt.Scan(&usr)
		fmt.Println()
		fmt.Println("Senha:")
		fmt.Scan(&pass)
		users, erro := os.Open("users.txt")
		secrets, erro := os.Open("secrets.txt")
		if erro != nil {
			fmt.Println("Erro: usuário não encontrado, tente novamente...")
			tentativas++
		}
		if tentativas == 3 {
			fmt.Println("Erro: excesso de tentativas para acessar o programa, forçando encerramento...")
			os.Exit(1)
		} else {
			usersReader := bufio.NewReader(users)
			// secretsReader := bufio.NewReader(secrets)
			var users []string
			for {
				linhaUsuarios, err := usersReader.ReadString('\n')
				linhaSecrets, err := bufio.NewReader(secrets).ReadString('\n')
				linhaUsuarios = strings.TrimSpace(linhaUsuarios)
				users = append(users, linhaUsuarios)

				if err != io.EOF {
					break
				}
			}
			for i := 0; i < len(users); i++ {
				if users[i] == usr {

				}
			}
			break
		}
	}
	for {
		entrada := input()
		switch entrada {
		case 1:
			salvarNovaSenha()
		case 2:
			mostrarSenhas()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Erro: comando não existe... saindo do programa.")
			os.Exit(1)
		}
	}

}

func mostrarSenhas() {
	panic("unimplemented")
}

func salvarNovaSenha() {
	panic("unimplemented")
}

func log(user string, pass string) {
	arquivo, erro := os.OpenFile(user+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if erro != nil {
		fmt.Println(erro)
	}
	arquivo.WriteString(user + " - " + pass + "\n")
	arquivo.Close()
}

func readTxt(user string) []string {
	res, _ := os.Open(user + ".txt")
	leitor := bufio.NewReader(res)
	var sites []string
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites
}

func input() int {
	fmt.Println("1 - Salvar nova senha")
	fmt.Println("2 - exibir senhas")
	fmt.Println("3 - sair")

	var cmd int
	fmt.Scan(&cmd)

	fmt.Printf("input recebido: %d \n", cmd)
	return cmd
}

func welcome() {
	var ent int
	fmt.Println("1 - criar usuário")
	fmt.Println("2 - entrar")
	fmt.Scan(&ent)
	if ent == 2 {
		auth()
	} else {
		newUser()
	}
}

func newUser() {
	panic("unimplemented")
}
