package utils

import (
	"compress/gzip"
	"os"
	"strings"
)

func CompressAndWriteToFile(lines []string, filename string) error {
	// Add ".gz" extension if not present
	if !strings.HasSuffix(filename, ".gz") {
		filename += ".gz"
	}

	// Create a gzip writer
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Write each line to the gzip writer
	for _, line := range lines {
		_, err := gzipWriter.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}