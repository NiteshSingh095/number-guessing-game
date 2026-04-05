package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/// ReadUserInput prompts the user for input and returns the input as an integer. It handles invalid inputs and continues to prompt until a valid number is entered.
func ReadUserInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		return num
	}
}

/// ReadUserInputString prompts the user for input and returns the input as a string. It handles empty inputs and continues to prompt until a valid string is entered.
func ReadUserInputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Invalid input. Please enter a valid string.")
			continue
		}

		return input
	}
}