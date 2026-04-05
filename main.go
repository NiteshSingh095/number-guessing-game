package main

import "fmt"

func main() {
	welcomeAndInfoMessage()
	var choice int

	for {
		_, err := fmt.Scanln(&choice)
		if err != nil {
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
}

/// getDifficultyRange returns the difficulty level and the number of attempts based on the user's choice.
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

/// welcomeAndInfoMessage displays the welcome message and game information to the user.
func welcomeAndInfoMessage() {

	fmt.Println("Welcome to Number Guessing Game 🚀")
	fmt.Println("\nEnter difficulty:\n1. Easy\n2. Medium\n3. Hard")
	fmt.Print("\nYour choice: ")
}
