package main

import "fmt"

//Crie um programa que receba o número do mês(1 a 12) e exiba a estaução do ano correspondente, considerando o hemisfério sul (Exemplo: (12,1,2) verão)

func main() {
	var mes int
	fmt.Print("Digite o número do mês (1 a 12): ")
	fmt.Scan(&mes)

	switch mes {
	case 12, 1, 2:
		fmt.Println("Verão")
	case 3, 4, 5:
		fmt.Println("Outono")
	case 6, 7, 8:
		fmt.Println("Inverno")
	case 9, 10, 11:
		fmt.Println("Primavera")
	default:
		fmt.Println("Mês inválido.")
	}
}