package main

import (
	"fmt"
	"math/rand"
	"time"
)

var chance int = 1

//create random number
func numberGenerator() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

// função que recebe input do número para adivinhar
func readInput() int {
	var input int
	fmt.Printf("Entre com sua aposta: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Deu xabu, tenta de novo!")
	} else {
		fmt.Printf("Sua aposta foi no número: %v\n\n", input)
	}
	return input
}

func isMatching(g, s int) bool {
	if g < s {
		fmt.Printf("Você apostou em um número MENOR que o número sorteado!\nTente de novo!\n\n")
		return false

	} else if g > s {
		fmt.Printf("Você apostou em um número MAIOR que o número sorteado\nTente de novo!\n\n")
		return false
	} else {
		fmt.Printf("Parabéns! Você acertou o número sorteado!\n\n")
		return true
	}
}

func main() {

	//cabeçalho
	fmt.Printf("\nBEM VINDO A LOTERIA DOS GURI\n\n")
	fmt.Printf("Um número aleatório de 0-10 foi gerado e você tem que adivinhar qual é!\nVocê tem TRÊS tentativas!\n\n")

	//generate random number
	secret := numberGenerator()

	guessed := readInput()
	acertou := isMatching(guessed, secret)
	var chance int = 1
	//make comparison

	for acertou != true && chance < 3 {
		guessed = readInput()
		acertou = isMatching(guessed, secret)
		chance++
	}

	if chance == 3 && acertou != true {
		fmt.Printf("Suas chances acabaram!!! Mais sorte na próxima!\n\n")

	}

}
