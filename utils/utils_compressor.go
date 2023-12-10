package utils

import (
	"compress/gzip"
	"fmt"
	"os"
	"strings"
)

func CompressAndWriteToFile(lines []string, filename string) error {
	// Create a gzip writer
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Join the lines into a single string
	content := strings.Join(lines, "\n")

	// Write the compressed content to the gzip writer
	_, err = gzipWriter.Write([]byte(content))
	if err != nil {
		return err
	}

	fmt.Println("Content compressed and written to", filename)

	return nil
}