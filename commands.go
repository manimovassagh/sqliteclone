package lite

import (
	"fmt"
	"strconv"
	"strings"
)

// handleSQLCommand processes SQL-like commands
func HandleSQLCommand(cmd string) {
	// Trim any leading/trailing spaces from the input
	cmd = strings.TrimSpace(cmd)

	// Split the command into parts based on spaces
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}

	// Handle different commands
	switch strings.ToUpper(parts[0]) {
	case "INSERT":
		// Check if the INSERT command is correctly formatted
		if len(parts) < 4 || parts[1] != "INTO" || parts[2] != "users" {
			fmt.Println("Invalid INSERT syntax. Use: INSERT INTO users <id> <name>")
			return
		}

		// Parse the ID and the name
		id, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Println("Invalid ID format.")
			return
		}

		name := strings.Join(parts[4:], " ")
		InsertRow(id, name)

	case "SELECT":
		// Check if the SELECT command is correctly formatted
		if len(parts) != 4 || parts[1] != "*" || parts[2] != "FROM" || parts[3] != "users" {
			fmt.Println("Invalid SELECT syntax. Use: SELECT * FROM users")
			return
		}

		// If the command is valid, proceed with SELECT operation
		SelectRows()

	default:
		// Handle unrecognized commands
		fmt.Println("Unrecognized command:", cmd)
	}
}