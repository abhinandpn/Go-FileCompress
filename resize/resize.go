package resize

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// ResizeAndSave resizes an image and saves it inside a folder named after the image
func ResizeAndSave(inputPath string, outputBaseDir string) error {
	// Open the original image file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Decode the image and detect format
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Extract filename without extension
	fileName := filepath.Base(inputPath)
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	// Create the output directory using the image name
	outputDir := filepath.Join(outputBaseDir, fileNameWithoutExt)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Define output files and sizes
	outputFiles := []struct {
		Name   string
		Width  uint
		Height uint
	}{
		{"standard", 0, 0},   // Full resolution
		{"medium", 800, 600}, // Medium resolution
		{"small", 300, 200},  // Thumbnail resolution
	}

	// Process and save images inside the folder
	for _, file := range outputFiles {
		// Set the correct file extension based on image format
		outputPath := filepath.Join(outputDir, file.Name+"."+format)
		if err := saveResizedImage(img, outputPath, file.Width, file.Height, format); err != nil {
			fmt.Println("Error resizing:", err)
		} else {
			fmt.Println("Saved:", outputPath)
		}
	}
	return nil
}

// saveResizedImage resizes and saves the image in the correct format
func saveResizedImage(img image.Image, outputPath string, width, height uint, format string) error {
	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the image based on format
	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(outFile, resizedImg, &jpeg.Options{Quality: 80})
	case "png":
		return png.Encode(outFile, resizedImg)
	case "gif":
		return gif.Encode(outFile, resizedImg, nil)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}
}
