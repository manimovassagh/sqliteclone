package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var dbFile string // Holds the name of the database file
var dbMutex sync.Mutex
var configFile = "config.txt" // File where the last used dbFile name is stored

// initDB initializes the database file
func InitDB() {
	// Check if config file exists, if so, use the last database file name from it
	if _, err := os.Stat(configFile); err == nil {
		// Read the config file to get the last used dbFile name
		file, err := os.Open(configFile)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		defer file.Close()

		// Read the database file name from the config file
		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			dbFile = scanner.Text()
		}
	}

	// If dbFile is still empty, ask the user for a new database file name
	if dbFile == "" {
		fmt.Print("Enter the database file name (without extension): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()

		// Append ".db" if no extension is provided
		if !strings.HasSuffix(userInput, ".db") {
			userInput += ".db"
		}

		// Set the dbFile variable to the user's input
		dbFile = userInput

		// Save the dbFile name to the config file
		SaveConfigFile(dbFile)
	}

	// Check if the database file exists, if not, create it
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Println("Database file doesn't exist, creating a new one.")
		CreateDatabaseFile()
	} else {
		// If the file exists, print that it's using the existing file
		fmt.Println("Database file exists, using existing file:", dbFile)
	}
}

// createDatabaseFile creates a new database file
func CreateDatabaseFile() {
	// Open the database file for writing
	file, err := os.Create(dbFile)
	if err != nil {
		fmt.Println("Error creating database file:", err)
		return
	}
	defer file.Close()

	// Initialize the database with headers
	file.WriteString("ID | Name\n")
}

// InsertRow inserts a new row into the database
func InsertRow(id int, name string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// Open the database file for appending
	file, err := os.OpenFile(dbFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening database file:", err)
		return
	}
	defer file.Close()

	// Write the new row to the file
	fmt.Fprintf(file, "%d | %s\n", id, name)
	fmt.Println("Row inserted successfully!")
}

// SelectRows selects and prints all rows from the database
func SelectRows() {
	// Open the database file for reading
	file, err := os.Open(dbFile)
	if err != nil {
		fmt.Println("Error opening database file:", err)
		return
	}
	defer file.Close()

	// Print the header
	fmt.Println("ID | Name")

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip the header line
		if line == "ID | Name" {
			continue
		}

		// Split the line into ID and Name
		parts := strings.Split(line, " | ")
		if len(parts) == 2 {
			fmt.Println(parts[0], "|", parts[1])
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from database file:", err)
	}
}

// saveConfigFile saves the database file name to a config file
func SaveConfigFile(dbName string) {
	// Open or create the config file to save the database name
	file, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return
	}
	defer file.Close()

	// Write the database name to the config file
	file.WriteString(dbName + "\n")
}
