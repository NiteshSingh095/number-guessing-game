package game

import (
	"fmt"
	"time"
	"math/rand"
	"number-guessing-game/utils"
	"number-guessing-game/internal/storage"
)


/// playGame is the main game loop that handles user guesses and provides feedback until the user either guesses correctly or exhausts all attempts.
func PlayGame(attempts int, randomNumber int, difficulty string) {
	guessCorrectly := false
	start := time.Now()
	cnt := 0
	wrongAttempts := 0
	hs, error := storage.LoadHighScore()
	if error != nil {
		fmt.Printf("Error loading high score: %v\n", error)
		return
	}

	for i := 0; i < attempts; i++ {
		fmt.Printf("Attempt %d/%d: Enter your guess: ", i+1, attempts)
		guess := input.ReadUserInput("")
		cnt++;
		if guess == randomNumber {
			fmt.Println("Congratulations! You've guessed the number correctly!")
			guessCorrectly = true
			break
		} else if guess < randomNumber {
			fmt.Println("Too low! Try again.")
			wrongAttempts++
		} else {
			fmt.Println("Too high! Try again.")
			wrongAttempts++
		}

		if wrongAttempts % 2 == 0 {
			hintSystem(randomNumber)
		}
	}

	if !guessCorrectly {
		fmt.Println("Sorry, you've used all your attempts!")
		fmt.Printf("Game Over! The correct number was %d.\n", randomNumber)
	} else {
		elapsed := time.Since(start)
		storage.UpdateHighScore(&hs, difficulty, cnt)
		err := storage.SaveHighScore(hs)
		if err != nil {
			fmt.Printf("Error saving high score: %v\n", err)
		}
		fmt.Printf("You guessed the number in %d attempts and %s!\n", cnt, elapsed.Round(time.Second))
	}
}

func hintSystem(num int) {
	fmt.Println("Here's a hint to help you out!. If you want to use hint, press yes (y/Y) or press no (n/N) to skip hint: ")
	choice := input.ReadUserInputString("")
	if choice != "y" && choice != "Y" {
		fmt.Println("No hint will be provided. Good luck!")
		return
	}
	randHintNum := GenerateRandomNumber(20)
	fmt.Printf("Hint: The number is between %d and %d :", (num-randHintNum)/2, num+randHintNum)
}



// / generateRandomNumber generates a random number between 1 and 100
func GenerateRandomNumber(num int) int {
	randomNumber := rand.Intn(num) + 1
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