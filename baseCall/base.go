package basecall

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetImage function to get user input correctly
func GetImage() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the full path of the image: ")

	// Read user input including spaces
	inputPath, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	inputPath = strings.TrimSpace(inputPath) // Trim newline or spaces
	return inputPath, nil
}

// Validation function to check if the file exists
func Validation(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist! Check the path and try again.")
		return err
	} else if err != nil {
		fmt.Println("Error: Unable to access file -", err)
		return err
	}
	return nil
}
