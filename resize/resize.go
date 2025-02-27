package resize

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"

	"github.com/nfnt/resize"
)

func init() {
	// Register additional image formats
	image.RegisterFormat("jpeg", "ÿØÿ", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "\x89PNG", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "GIF", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("bmp", "BM", bmp.Decode, bmp.DecodeConfig)
	image.RegisterFormat("tiff", "II*\x00", tiff.Decode, tiff.DecodeConfig)
	image.RegisterFormat("webp", "RIFF????WEBPVP", webp.Decode, webp.DecodeConfig)
}

// ResizeAndSave resizes an image and saves it inside a folder named after the image
func ResizeAndSave(inputPath string, outputBaseDir string, targetSizeKB int) error {
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

	// Ensure format is normalized (fixes JPEG vs JPG issues)
	if format == "jpeg" {
		format = "jpg"
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
		{"standard", uint(img.Bounds().Dx()), uint(img.Bounds().Dy())}, // Full resolution
		{"medium", 800, 600}, // Medium resolution
		{"small", 300, 200},  // Thumbnail resolution
	}

	// Process and save images inside the folder
	for _, file := range outputFiles {
		// Set the correct file extension based on image format
		outputPath := filepath.Join(outputDir, file.Name+"."+format)
		if err := saveResizedImage(img, outputPath, file.Width, file.Height, format, targetSizeKB); err != nil {
			fmt.Println("Error resizing:", err)
		} else {
			fmt.Println("Saved:", outputPath)
		}
	}
	return nil
}

// saveResizedImage resizes and saves the image in the correct format while limiting file size (KB)
func saveResizedImage(img image.Image, outputPath string, width, height uint, format string, targetSizeKB int) error {
	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Use a buffer to dynamically adjust quality
	var buffer bytes.Buffer

	switch format {
	case "jpg":
		quality := 90 // Start with high quality
		for {
			buffer.Reset()
			jpeg.Encode(&buffer, resizedImg, &jpeg.Options{Quality: quality})

			// Check file size
			if buffer.Len() <= targetSizeKB*1024 || quality <= 10 {
				break
			}
			quality -= 5 // Reduce quality if the file is too big
		}
		_, err = outFile.Write(buffer.Bytes())

	case "png":
		// PNG compression is lossless; file size is hard to control
		png.Encode(&buffer, resizedImg)
		if buffer.Len() > targetSizeKB*1024 {
			fmt.Println("Warning: PNG file size exceeds target limit.")
		}
		_, err = outFile.Write(buffer.Bytes())

	case "gif":
		// Reduce colors for GIF optimization
		err = gif.Encode(outFile, resizedImg, &gif.Options{NumColors: 256})

	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	return err
}
