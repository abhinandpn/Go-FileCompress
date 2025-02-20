package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/abhinandpn/Go-FileCompress/resize"
)

func main() {
	// Ask the user for the image path
	var inputPath string
	fmt.Print("Enter the full path of the image: ")
	fmt.Scanln(&inputPath)

	// Trim spaces from the input
	inputPath = strings.TrimSpace(inputPath)

	// Validate if the file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist! Check the path and try again.")
		return
	}

	// Set the base output directory
	outputBaseDir := "resize"

	// Call the function to resize and save images
	err := resize.ResizeAndSave(inputPath, outputBaseDir)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Images successfully saved in:", outputBaseDir)
	}
}
