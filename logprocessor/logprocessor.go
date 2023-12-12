package logprocessor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/5ud03r5/go-siem-agent/utils"
)

func RunProcessors(config *[]utils.ConfigEntry, wg *sync.WaitGroup) {
    for _, entry := range *config {
        wg.Add(1)
		directoryName := "backlog/" + entry.Name
		utils.CreateDirectory(directoryName)
        go func(name string, filePath string, format string, backlogSize int64) {
            defer wg.Done()
            processLog(name, filePath, format, backlogSize)
        }(entry.Name, entry.FilePath, entry.Format, entry.BacklogSize)
    }
}

func processLog(name string, filePath string, format string, backlogSize int64) {
	var position int64
	position = 0

	for {
		directoryName := "backlog/" + name
		size, err := utils.GetDirectorySize(directoryName)

		if err != nil {
			panic(err)
		}

		if size > backlogSize {
			fmt.Printf("%v: Directory size reached treshold of %v\n", directoryName, backlogSize)
		}
		returnedPosition, lineCount, processedLines, err := processNewLines(filePath, position)
		if err != nil {
			panic(err)
		}
		position = returnedPosition
		unixTimestamp := time.Now().Unix()
		unixTimestampString := strconv.FormatInt(unixTimestamp, 10)
		filename := directoryName + "/" + name + "_" + unixTimestampString + "_" + strconv.FormatInt(position, 10)
		if lineCount > 0 {
			utils.CompressAndWriteToFile(processedLines, filename)
		}

		duration := 5 * time.Second
		time.Sleep(duration)
	}
}

func processNewLines(fileName string, lastPosition int64) (int64, int, []string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, 0, nil, err
	}
	fileSize := fileInfo.Size()

	// If the last position is greater than the file size, set it to 0
	if lastPosition > fileSize {
		lastPosition = 0
	}

	// Set the file offset to the last position
	file.Seek(lastPosition, 0)

	scanner := bufio.NewScanner(file)
	lineCount := 0
	var processedLines []string
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line (replace this with your actual processing logic)
		// fmt.Println("Processing line:", line)
		processedLines = append(processedLines, line)
		lineCount++
	}

	// Get the current offset after processing
	currentPosition, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, 0, nil, err
	}

	return currentPosition, lineCount, processedLines, nil
}