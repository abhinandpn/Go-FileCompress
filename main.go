package main

import (
	"fmt"

	basecall "github.com/abhinandpn/Go-FileCompress/baseCall"
	"github.com/abhinandpn/Go-FileCompress/resize"
)

func main() {

	// Set the base output directory
	outputBaseDir := "resize"

	inputPath, err := basecall.GetImage()
	if err != nil {
		fmt.Println("Error:", err)
		return // Exit the program
	}

	err = basecall.Validation(inputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return // Exit the program
	}

	// Call the function to resize and save images
	err = resize.ResizeAndSave(inputPath, outputBaseDir)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Images successfully saved in:", outputBaseDir)
	}
}
