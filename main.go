package main

import (
	"hangman/pkg/screen"
	"hangman/pkg/words"
	"os"
	"strings"
)

const max_misses = 10
const MASKING_CHAR = "_"

func main() {
	gameLoop()
}

func gameLoop() {
	word := fetchWord()
	solved := false
	remainingTries := max_misses
	known := []string{}
	for remainingTries > 1 {

		screen.PrintIntro()
		maskedWord := getWordMask(word, known)
		screen.PrintWord(maskedWord)
		screen.PrintTry(remainingTries - 1)
		screen.PrintPrompt()

		char := getNextLetter()

		if !contains(word, char) {
			remainingTries--
		}

		known = append(known, char)
		solved = validSolution(word, known)
		if solved {
			screen.PrintWin(word)
			break
		}

	}
	if !solved {
		screen.PrintLose(word)
	}
}

func validSolution(word string, known []string) bool {
	solution := getWordMask(word, known)
	return !strings.Contains(solution, MASKING_CHAR)
}

func getWordMask(word string, known []string) string {

	val := strings.Split(word, "")
	for i, _ := range word {
		val[i] = MASKING_CHAR
	}

	for i, char := range word {
		for _, k := range known {
			if string(rune(char)) == k {
				val[i] = k
			}
		}
	}
	return strings.Join(val, "")
}

func contains(word, char string) bool {
	return strings.Contains(word, char)
}

func getNextLetter() string {
	var b []byte = make([]byte, 3)

	os.Stdin.Read(b)

	return strings.ToLower(string(b[0]))
}

func fetchWord() string {
	return words.GetRandomWord()
}
