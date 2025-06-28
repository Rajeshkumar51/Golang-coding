package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
	var wg sync.WaitGroup
	errorChan := make(chan string)
	writerDone := make(chan error)
	go func() {
		err := writeErrorsToFile(outputFile, errorChan)
		writerDone <- err
	}()
	for _, file := range inputFiles {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			if err := processFile(filePath, errorChan); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to process %s: %v\n", filePath, err)
			}
		}(file)
	}
	go func() {
		wg.Wait()
		close(errorChan)
	}()
	return <-writerDone
}
func processFile(filePath string, errorChan chan<- string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorChan <- line
		}
	}
	return scanner.Err()
}
func writeErrorsToFile(outputFile string, errorChan <-chan string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for line := range errorChan {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
