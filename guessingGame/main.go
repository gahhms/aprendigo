package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
		fmt.Printf("Sua aposta foi no número: %v\n", input)
	}
	return input
}

func isMatching(g, s int) bool {
	if g < s {
		fmt.Println("Sua aposta é MENOR que o número sorteado!")
		return false

	} else if g > s {
		fmt.Println("Sua aposta é MAIOR que o número sorteado!")
		return false
	} else {
		fmt.Println("Parabéns! Você acertou o número sorteado!")
		return true
	}
}
func main() {

	//cabeçalho
	fmt.Printf("\nBEM VINDO A LOTERIA DOS GURI\n\n")
	fmt.Printf("Um número aleatório de 0-10 foi gerado e você tem que adivinhar qual é!\n\n")

	//generate random number
	secret := numberGenerator()

	// get user input
	guessed := readInput()
	acertou := isMatching(guessed, secret)
	//make comparison
	for acertou != true {
		guessed = readInput()
		acertou = isMatching(guessed, secret)
	}

}
