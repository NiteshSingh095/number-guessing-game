package game

import (
	"fmt"
	"math/rand"
	"number-guessing-game/utils"
)


/// playGame is the main game loop that handles user guesses and provides feedback until the user either guesses correctly or exhausts all attempts.
func PlayGame(attempts int, randomNumber int) {
	guessCorrectly := false

	for i := 0; i < attempts; i++ {
		fmt.Printf("Attempt %d/%d: Enter your guess: ", i+1, attempts)
		guess := input.ReadUserInput("")

		if guess == randomNumber {
			fmt.Println("Congratulations! You've guessed the number correctly!")
			guessCorrectly = true
			break
		} else if guess < randomNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}

	if !guessCorrectly {
		fmt.Println("Sorry, you've used all your attempts!")
		fmt.Printf("Game Over! The correct number was %d.\n", randomNumber)
	}
}



// / generateRandomNumber generates a random number between 1 and 100
func GenerateRandomNumber() int {
	randomNumber := rand.Intn(100) + 1
	return randomNumber
}

// / getDifficultyRange returns the difficulty level and the number of attempts based on the user's choice.
func GetDifficultyRange(choice int) (string, int) {
	switch choice {
	case 1:
		return "Easy", 10

	case 2:
		return "Medium", 5

	case 3:
		return "Hard", 3

	default:
		return "Invalid", 0
	}
}