package main

import (
	"fmt"
	"math/rand"
)

func main() {
	welcomeAndInfoMessage()
	var choice int

	for {
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Error reading input. Please enter a valid number.")
			continue
		}
		break
	}

	level, attempts := getDifficultyRange(choice)
	if attempts == 0 {
		fmt.Println("Invalid choice. Please select a valid difficulty level.")
		return
	}

	fmt.Printf("You have selected %s difficulty. You will have %d attempts to guess the number.\n", level, attempts)

	randomNumber := generateRandomNumber()

	gameCore(attempts, randomNumber)
}

/// gameCore is the main game loop that handles user guesses and provides feedback until the user either guesses correctly or exhausts all attempts.
func gameCore(attempts int, randomNumber int) {
	var guess int
	guessCorrectly := false

	for i := 0; i < attempts; i++ {
		fmt.Printf("Attempt %d: Enter your guess: ", i+1)
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Error reading input. Please enter a valid number.")
			i--
			continue
		}

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
func generateRandomNumber() int {
	randomNumber := rand.Intn(100) + 1
	return randomNumber
}

// / getDifficultyRange returns the difficulty level and the number of attempts based on the user's choice.
func getDifficultyRange(choice int) (string, int) {
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

// / welcomeAndInfoMessage displays the welcome message and game information to the user.
func welcomeAndInfoMessage() {

	fmt.Println("Welcome to Number Guessing Game 🚀")
	fmt.Println("\nEnter difficulty:\n1. Easy\n2. Medium\n3. Hard")
	fmt.Print("\nYour choice: ")
}
