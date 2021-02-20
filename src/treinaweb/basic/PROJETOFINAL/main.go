package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ARQUIVO : constante criada para o arquivo agenda.txt
const ARQUIVO string = "agenda.txt"

// ConversivelParaString : interface que especifica como um item pode ser convertido para uma string
type ConversivelParaString interface {
	toString() string
}

// Contato : estrutura que representa um contato
type Contato struct {
	nome         string
	formaContato string
	valorContato string
}

// toString : muda para string
func (contato *Contato) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contato.nome, contato.formaContato, contato.valorContato)
}

// GerenciadorContatos : componente resposável por gerencia contatos
type GerenciadorContatos struct{}

// carregarContatos para carregar os contatos
func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error) {
	contatos := make([]Contato, 0)
	if _, e := os.Stat(ARQUIVO); !os.IsNotExist(e) {
		arquivo, err := os.Open(ARQUIVO)
		if err != nil {
			return contatos, err
		}
		defer arquivo.Close()
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhaContato := scanner.Text()
			infoContato := strings.Split(linhaContato, "|")
			contatos = append(contatos, Contato{nome: infoContato[0], formaContato: infoContato[1], valorContato: infoContato[2]})
		}
	}
	return contatos, nil
}

// salvarContato : salva os contatos no arquivo
func (gerenciador *GerenciadorContatos) salvarContato(contato ConversivelParaString) error {
	arquivo, err := os.OpenFile(ARQUIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	if _, e := arquivo.WriteString(contato.toString()); e != nil {
		return e
	}
	return nil
}

func main() {
	fmt.Println("")
	fmt.Println("  Projeto Final")
	fmt.Println("  Agenda de Contatos")
	fmt.Println("")
	fmt.Println(".........................: ")
	gerenciador := GerenciadorContatos{}
	opcao := 0
	for true {
		fmt.Println(".........................: ")
		fmt.Println(" O que você deseja fazer? ")
		fmt.Println(".........................: ")
		fmt.Println(" 1 - Listar os contatos")
		fmt.Println(" 2 - Criar novo contatos")
		fmt.Println(" 3 - Sair da agenda ")
		fmt.Print(".........................: ")
		fmt.Scanf("%d", &opcao)

		switch opcao {
		case 1:
			fmt.Println("")
			listarContatos(&gerenciador)
			fmt.Println("")
		case 2:
			fmt.Println("")
			criarNovoContato(&gerenciador)
			fmt.Println("")
		}
		if opcao == 3 {
			break
		}
	}
	fmt.Println("Bye ;)")

}

// listarContatos : listar contatos cadastrados na agenda
func listarContatos(gerenciador *GerenciadorContatos) {
	contatos, err := gerenciador.carregarContatos()
	if err != nil {
		fmt.Printf("Não foi possível carregar os contatos: %s", err)
	} else {
		fmt.Println("Lista de contatos: ")
		for _, contato := range contatos {
			fmt.Printf(" - %s, %s: %s \n", contato.nome, contato.formaContato, contato.valorContato)
		}
	}
}

// criarNovoContato : cria um novo contato
func criarNovoContato(gerenciador *GerenciadorContatos) {
	novoContato := Contato{}
	fmt.Print("Nome do contato: ")
	fmt.Scanf("%s", &novoContato.nome)
	fmt.Print("Tipo do contato: ")
	fmt.Scanf("%s", &novoContato.formaContato)
	fmt.Print("Contato: ")
	fmt.Scanf("%s", &novoContato.valorContato)
	err := gerenciador.salvarContato(&novoContato)
	if err != nil {
		fmt.Printf("Houve um erro ao salvar o contato: %s \n", err)
	}
}
