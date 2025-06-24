package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
	var wg sync.WaitGroup
	errorChan := make(chan string)
	done := make(chan error)
	go func() {
		done <- writeErrors(outputFile, errorChan)
	}()
	for _, file := range inputFiles {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := extractErrors(filename, errorChan)
			if err != nil {
				log.Printf("Error processing %s: %v", filename, err)
			}
		}(file)
	}
	wg.Wait()
	close(errorChan)
	return <-done
}
func extractErrors(filename string, errorChan chan<- string) error {
	file, err := os.Open(filename)
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
func writeErrors(outputFile string, errorChan <-chan string) error {
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func main() {
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	err := ProcessLogs(inputFiles, "errors.log")
	if err != nil {
		log.Fatal("Failed to process logs:", err)
	} else {
		fmt.Println("Error lines extracted successfully.")
	}
}
