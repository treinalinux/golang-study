package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Mapas são muito importantes para relacionamento:")
	mapaCursos := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	curso := ""
	for curso != "q" {
		var cargaHoraria int
		fmt.Print("Digite um curso ou digite 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			fmt.Print("Digite a carga horária do curso: ")
			fmt.Scanf("%d", &cargaHoraria)
			mapaCursos[curso] = cargaHoraria
		}
	}
	fmt.Println("Cursos registrados: ")
	// Perceba que cargaHoraria aqui está em outro escopo:
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf("- %s: %dh \n", nomeCurso, cargaHoraria)
	}
	curso = ""
	for curso != "q" {
		fmt.Print("Digite o nome do curso a ser excluído ou digite 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			cargaHoraria, cursoExiste := mapaCursos[curso]
			if cursoExiste {
				delete(mapaCursos, curso)
				fmt.Printf("O curso %s com %dh foi excluído com sucesso: \n", curso, cargaHoraria)
			} else {
				fmt.Println("O curso digitado não existe.")
			}
		}
	}
	fmt.Println("Cursos registradas: ")
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf(" - %s: %dh \n", nomeCurso, cargaHoraria)
	}
}
