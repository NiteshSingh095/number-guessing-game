package cmd

import (
	"fmt"
	"number-guessing-game/game"
	"number-guessing-game/utils"
)

const maxNumber = 100

func Execute() {
	for {

		welcomeAndInfoMessage()

		choice := input.ReadUserInput("Please enter your choice (1, 2, or 3):")

		level, attempts := game.GetDifficultyRange(choice)
		if attempts == 0 {
			fmt.Println("Invalid choice. Please select a valid difficulty level.")
			return
		}

		fmt.Printf("You have selected %s difficulty. You will have %d attempts to guess the number.\n", level, attempts)

		randomNumber := game.GenerateRandomNumber(maxNumber)

		game.PlayGame(attempts, randomNumber, level)

		fmt.Println("\nDo you want to play again? (y/n): ")
		playAgain := input.ReadUserInputString("")
		if playAgain != "y" && playAgain != "Y" {
			fmt.Println("Thank you for playing! Goodbye!")
			break
		}
	}
}

// / welcomeAndInfoMessage displays the welcome message and game information to the user.
func welcomeAndInfoMessage() {

	fmt.Println("Welcome to Number Guessing Game 🚀")
	fmt.Println("\nEnter difficulty:\n1. Easy\n2. Medium\n3. Hard")
	fmt.Print("\nYour choice: ")
}
