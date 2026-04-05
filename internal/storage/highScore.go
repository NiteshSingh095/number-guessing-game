package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"number-guessing-game/model"
)

/// SaveHighScore saves the high score to a JSON file. It takes the file name and the high score data as parameters and returns an error if any occurs during the file writing process.
func LoadHighScore() (model.HighScore, error) {
	var hs model.HighScore

	file, err := os.Open("highScore.json")
	if err != nil {
		if os.IsNotExist(err) {
			return hs, nil // Return empty high score if file does not exist
		}
		return hs, fmt.Errorf("error opening high score file: %v", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&hs)
	if err != nil {
		return hs, fmt.Errorf("error decoding high score file: %v", err)
	}

	return hs, nil
}

/// SaveHighScore saves the high score to a JSON file. It takes the file name and the high score data as parameters and returns an error if any occurs during the file writing process.
func SaveHighScore(hs model.HighScore) error {
	file, err := os.Create("highScore.json")
	if err != nil {
		return fmt.Errorf("error creating high score file: %v", err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(hs)
	if err != nil {
		return fmt.Errorf("error encoding high score to file: %v", err)
	}

	return nil
}

/// UpdateHighScore updates the high score based on the difficulty level and the number of attempts taken. 
// It takes the current high score, difficulty level, and attempts as parameters and 
// updates the high score if the new score is better than the existing one.
func UpdateHighScore(hs *model.HighScore, difficulty string, attempts int) error {
	switch difficulty {
	case "Easy":
		if hs.Easy == 0 || attempts < hs.Easy {
			hs.Easy = attempts
		}
	case "Medium":
		if hs.Medium == 0 || attempts < hs.Medium {
			hs.Medium = attempts
		}
	case "Hard":
		if hs.Hard == 0 || attempts < hs.Hard {
			hs.Hard = attempts
		}
	}
	return nil
}