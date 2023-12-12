package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectory(directoryPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, so create it
		err := os.Mkdir(directoryPath, 0755)
		if err != nil {
			return err
		}
		fmt.Printf("Directory %s created successfully\n", directoryPath)
	} else if err != nil {
		return err
	}
	return nil
}

func GetDirectorySize(directoryPath string) (int64, error) {
	var size int64

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}