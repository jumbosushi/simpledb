package main

import (
	"bufio"
	"fmt"
	"os"
)

// Version defines current version
var Version = "0.0.1"

// Database is data type for database
type Database struct {
	Name   string
	Tables []Table
}

// Table is data type for table
type Table struct {
	Name string
}

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
	switch input {
	case "quit\n":
		os.Exit(0)
	default:
		fmt.Println(input)
	}
	Prompt()
}

// Intro returns introstring
func Intro() string {
	return fmt.Sprintf("simpledb (%s)", Version)
}
