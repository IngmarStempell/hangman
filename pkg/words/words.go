package words

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetRandomWord() string {
	rand.Seed(time.Now().UnixNano())

	lines := readLines()
	length := len(lines)
	println(length)
	lineIndex := rand.Intn(length)
	println(lineIndex)
	return lines[lineIndex]
}

func readLines() []string {
	lines := []string{}
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
