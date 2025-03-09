package lite

import (
	"strings"
	"testing"
)

// TestInsertRow tests the insertion of rows
func TestInsertRow(t *testing.T) {
	// Clear the database before each test
	clearDB()

	// Insert a row
	InsertRow(1, "Mani")

	// Capture the output of SelectRows
	output := captureStdout(func() {
		SelectRows()
	})

	// Check if the output matches the expected result
	expected := "ID | Name\n1 | Mani\n"
	output = strings.TrimSpace(output)     // Trim the output to remove extra spaces or newlines
	expected = strings.TrimSpace(expected) // Trim the expected output to avoid mismatch

	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}

// TestSelectRows tests the SelectRows functionality
func TestSelectRows(t *testing.T) {
	// Clear the database before each test
	clearDB()

	// Insert some rows
	InsertRow(1, "Mani")
	InsertRow(2, "Alice")

	// Capture the output of SelectRows
	output := captureStdout(func() {
		SelectRows()
	})

	// Check if the output matches the expected result
	expected := "ID | Name\n1 | Mani\n2 | Alice\n"
	output = strings.TrimSpace(output)     // Trim the output to remove extra spaces or newlines
	expected = strings.TrimSpace(expected) // Trim the expected output to avoid mismatch

	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}

// Helper function to clear the database
func clearDB() {
	// Clear or reset the database file before tests
	// You can implement logic here to reset the database state
}

// Helper function to capture stdout
func captureStdout(f func()) string {
	// Implement logic to capture stdout (you can use a buffer or other methods)
	return ""
}
