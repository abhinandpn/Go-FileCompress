package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage resizes the image to the given width and height, and saves it
func ResizeImage(inputPath, outputPath string, width, height uint) error {
	// Open the original image file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Encode the resized image as JPEG with quality 80
	err = jpeg.Encode(outFile, resizedImg, &jpeg.Options{Quality: 80})
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	fmt.Println("Saved:", outputPath)
	return nil
}

func main() {
	// Input image path
	inputImage := "input.jpg"

	// Define output sizes
	outputFiles := []struct {
		Path   string
		Width  uint
		Height uint
	}{
		{"output_standard.jpg", 0, 0},   // Full resolution
		{"output_medium.jpg", 800, 600}, // Medium resolution
		{"output_small.jpg", 300, 200},  // Thumbnail resolution
	}

	// Process and save images
	for _, file := range outputFiles {
		err := ResizeImage(inputImage, file.Path, file.Width, file.Height)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
