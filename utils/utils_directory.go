package utils

import (
	"fmt"
	"os"
)

func CreateDirectory(directoryPath string) error {
	// Create the directory
	err := os.Mkdir(directoryPath, 0755)
	if err != nil {
		return err
	}

	fmt.Printf("Directory %s created successfully\n", directoryPath)
	return nil
}