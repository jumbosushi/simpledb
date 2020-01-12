package main

import (
	"bufio"
	"fmt"
	"os"
)

// Version defines current version
var Version = "0.0.1"

func main() {
	fmt.Println(Intro())
	fmt.Println("Type \"help\" for help")
	Prompt()
}

// Prompt asks prompt
func Prompt() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	input := string(sentence)
	if input == "quit\n" {
		os.Exit(0)
	} else {
		fmt.Println(input)
	}
	Prompt()
}

// Intro returns introstring
func Intro() string {
	return fmt.Sprintf("simpledb (%s)", Version)
}
