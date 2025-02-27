package main

import (
	"fmt"
	"log"
	"os"

	basecall "github.com/abhinandpn/Go-FileCompress/baseCall"
	"github.com/abhinandpn/Go-FileCompress/resize"
)

func main() {
	// Get user input for the image path
	inputPath, err := basecall.GetImage()
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	// Validate if the file exists
	if err := basecall.Validation(inputPath); err != nil {
		log.Fatal("File validation failed:", err)
	}

	// Ensure output directory exists
	outputBaseDir := "output/directory"
	if err := os.MkdirAll(outputBaseDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	// Set the target file size (in KB)
	targetSizeKB := 100

	// Resize and save the image
	if err := resize.ResizeAndSave(inputPath, outputBaseDir, targetSizeKB); err != nil {
		log.Fatal("Error resizing image:", err)
	}

	fmt.Println("Image compression completed successfully!")
}
