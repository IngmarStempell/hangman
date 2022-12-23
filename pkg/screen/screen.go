package screen

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
)

func pprint(word string) {
	pword := figure.NewFigure(word, "", true)
	pword.Print()
}

func pprintI(word string) {
	pword := figure.NewFigure(word, "isometric1", true)
	pword.Print()
}

func printGap() {
	for i := 0; i < 10; i++ {
		fmt.Println("")
	}
}

func PrintIntro() {
	printGap()
	printGap()
	printFile("intro.txt")

}

func PrintTry(try int) {
	printFile(fmt.Sprintf("%d-tries.txt", try))
}

func PrintWord(word string) {
	printGap()
	printGap()
	pprint(word)
	printGap()
}

func PrintLose(word string) {
	printGap()
	printGap()
	printFile("lose.txt")
	pprint("The Word was: ")
	pprint(word)
	printGap()
	printGap()
}
func PrintWin(word string) {
	printGap()
	printGap()
	printGap()
	printFile("win.txt")
	pprint("The Word was: ")
	pprint(word)
	printGap()
	printGap()
	printGap()
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char
	}
}

func PrintPrompt() {
	fmt.Println("Enter next letter and press enter")
}
