package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize the database
	initDB()

	// Command prompt
	fmt.Println("Welcome to GoSQLLite! Enter commands (type '.exit' to quit):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Prompt for user input
		fmt.Print("GoSQLLite> ")
		scanner.Scan()
		cmd := scanner.Text()

		// Exit the program if the user types .exit
		if cmd == ".exit" {
			fmt.Println("Exiting...")
			break
		}

		// Handle SQL-like commands
		handleSQLCommand(cmd)
	}
}
