# 🎯 Number Guessing Game (CLI - Go)

A fun and interactive **Command Line Number Guessing Game** built using **Go (Golang)**.
Test your luck, improve your guessing skills, and try to beat your high score!

---

## 🚀 Features

* 🎮 CLI-based interactive gameplay
* 🎚️ Multiple difficulty levels:

  * Easy (10 attempts)
  * Medium (5 attempts)
  * Hard (3 attempts)
* 🔁 Replay system (play multiple rounds)
* ⏱️ Timer tracking (measure how fast you win)
* 🧠 Hint system (helps when stuck)
* 🏆 High score tracking (persistent storage using JSON)
* ✅ Input validation and clean error handling
* 🧱 Modular project structure (production-style Go code)

---

## Project URL : https://roadmap.sh/projects/number-guessing-game

## 📁 Project Structure

```
number-guessing-game/
│
├── main.go
├── game/
│   └── game.go
├── utils/
│   └── input.go
├── internal/
│   └── storage/
│       └── highscore.go
├── highscore.json
└── go.mod
```

---

## 🛠️ Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/NiteshSingh095/number-guessing-game
cd number-guessing-game
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the game

```bash
go run main.go
```

---

## 🎮 How to Play

1. Start the game
2. Choose a difficulty level:

   * `1` → Easy
   * `2` → Medium
   * `3` → Hard
3. Guess the number between **1 and 100**
4. Get feedback after each guess:

   * 📉 Too low
   * 📈 Too high
5. Win before running out of attempts

---

## 🧠 Example Gameplay

```
Welcome to Number Guessing Game 🚀

Enter difficulty:
1. Easy
2. Medium
3. Hard

Your choice: 2

You have selected Medium difficulty.

Attempt 1/5: Enter your guess: 50
Too high!

Attempt 2/5: Enter your guess: 25
Too low!

Attempt 3/5: Enter your guess: 30
🎉 Congratulations! You guessed the number!

Time taken: 8 seconds
```

---

## 🏆 High Score System

* Best scores are stored in `highscore.json`
* Tracks lowest attempts per difficulty
* Automatically updates when a better score is achieved

---

## 🧩 Concepts Covered

This project demonstrates:

* Go Modules (`go.mod`)
* Package structuring
* CLI input handling (`bufio`, `os`)
* Random number generation (`math/rand`)
* Loops & control flow
* Error handling
* JSON encoding/decoding
* File I/O operations
* Clean architecture principles

---

## 💡 Future Improvements

* 🎨 Add colored CLI output
* 🌐 Convert to REST API
* 🖥️ Build GUI using Go or Flutter
* 📊 Add analytics/statistics tracking

---

## 👨‍💻 Author

Built with ❤️ using Go

---

## ⭐ Support

If you like this project:

* ⭐ Star the repository
* 🍴 Fork it
* 🚀 Enhance it further

---
