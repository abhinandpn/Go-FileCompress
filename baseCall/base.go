package basecall

import (
	"fmt"
	"os"
	"strings"
)

func GetImage() (string, error) {

	var inputPath string
	fmt.Print("Enter the full path of the image: ")
	fmt.Scanln(&inputPath)

	inputPath = strings.TrimSpace(inputPath)
	return inputPath, nil
}

func Validation(string)error {

	var inputPath string
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist! Check the path and try again.")
		return err
	}
	return nil

}
