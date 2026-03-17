package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 1. Get the full path in the Home Directory
	filePath, err := getTimestampPath()
	if err != nil {
		fmt.Printf("Error finding home directory: %v\n", err)
		return
	}

	// 2. Generate the data independently
	dataRow := generateRandomData()

	// 3. Pass the generated data to the append function
	err = appendToCSV(filePath, dataRow)
	if err != nil {
		fmt.Printf("Error updating CSV: %v\n", err)
		return
	}

	fmt.Printf("Data appended successfully to: %s\n", filePath)

	// 4 second delay before next run
	time.Sleep(4 * time.Second)

	// 5. Generate the data independently
	dataRow2 := generateRandomData()

	// 6. Pass the generated data to the append function
	err = appendToCSV(filePath, dataRow2)
	if err != nil {
		fmt.Printf("Error updating CSV: %v\n", err)
		return
	}

	fmt.Printf("Data appended successfully to: %s\n", filePath)
}

// getTimestampPath defines where the file lives
func getTimestampPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// Filename based on date so it appends to the same file throughout the day
	fileName := fmt.Sprintf("data_%s.csv", time.Now().Format("2006-01-02"))
	return filepath.Join(home, fileName), nil
}

// generateRandomData creates the timestamp and 4 random floats
func generateRandomData() []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	currentTime := time.Now().Format(time.RFC3339)
	row := []string{currentTime}

	for i := 0; i < 4; i++ {
		row = append(row, fmt.Sprintf("%.4f", r.Float64()*100))
	}
	return row
}

// appendToCSV now accepts the data as an input parameter
func appendToCSV(fullPath string, record []string) error {
	// Check if file exists to determine if we need a header
	_, err := os.Stat(fullPath)
	isNewFile := os.IsNotExist(err)

	// Open with Append and Create flags
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write Headers only for a brand new file
	if isNewFile {
		headers := []string{"Timestamp", "Val1", "Val2", "Val3", "Val4"}
		if err := writer.Write(headers); err != nil {
			return err
		}
	}

	// Write the record passed from main
	return writer.Write(record)
}
