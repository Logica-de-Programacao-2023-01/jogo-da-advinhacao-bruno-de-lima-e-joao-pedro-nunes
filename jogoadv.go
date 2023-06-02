package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	var playAgain string
	var totalAttempts [][]int

	for {
		answer := generateRandomNumber(1, 100)
		attempts := make([]int, 0)

		fmt.Print("Digite um número entre 1 e 100: ")
		var guess int
		fmt.Scanf("%d", &guess)
		attempts = append(attempts, guess)

		for guess != answer {
			if guess > answer {
				fmt.Println("O número é menor que", guess)
			} else {
				fmt.Println("O número é maior que", guess)
			}

			fmt.Print("Digite um número entre 1 e 100: ")
			fmt.Scanf("%d", &guess)
			attempts = append(attempts, guess)
		}

		fmt.Println("Parabéns, você acertou!")
		fmt.Println("Você utilizou", len(attempts), "tentativas.")

		totalAttempts = append(totalAttempts, attempts)

		fmt.Print("Você deseja jogar novamente? (s/n): ")
		fmt.Scanf("%s", &playAgain)

		if playAgain != "s" && playAgain != "S" {
			break
		}
	}

	for i, attempts := range totalAttempts {
		fmt.Printf("Você utilizou %d tentativas na %s jogada.\n", len(attempts), ordinalSuffix(i+1))
	}
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func ordinalSuffix(i int) string {
	if i%10 == 1 && i%100 != 11 {
		return fmt.Sprintf("%dst", i)
	}
	if i%10 == 2 && i%100 != 12 {
		return fmt.Sprintf("%dnd", i)
	}
	if i%10 == 3 && i%100 != 13 {
		return fmt.Sprintf("%drd", i)
	}
	return fmt.Sprintf("%dth", i)
}
