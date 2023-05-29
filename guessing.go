package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxNum = 100

func main() {
	secretNumber := generateSecretNumber()
	playGame(secretNumber)
}

// generateSecretNumber returns a random number within a specified range.
func generateSecretNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum)
}

// playGame executes the guessing game.
func playGame(secretNumber int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your guess")

	for {
		guess := readGuess(reader)
		if isCorrectGuess(guess, secretNumber) {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}

// readGuess prompts for and reads the player's guess.
func readGuess(reader *bufio.Reader) int {
	for {
		fmt.Println("Enter your guess:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}

		guess, err := strconv.Atoi(strings.Trim(input, "\r\n"))
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		fmt.Println("Your guess is", guess)
		return guess
	}
}

// isCorrectGuess checks if the guess matches the secret number and gives the player a hint.
func isCorrectGuess(guess, secretNumber int) bool {
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Please try again")
		return false
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Please try again")
		return false
	}

	return true
}
