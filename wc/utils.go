package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true // Target found
		}
	}
	return false // Target not found
}

func countWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Set scanner to split by words

	wordCount := 0

	// Count each word
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return wordCount, nil
}

func countChars(fileName string) (int, error) {

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}

	// Get the size in bytes
	size := fileInfo.Size()

	return int(size), nil

}

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Count each line
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return lineCount, nil
}

// parseArguments parses the command-line arguments and returns the desired field and the file to open.
func parseArguments(args []string) ([]string, string, error) {
	var desiredFlags []string
	var fileToOpen string = ""

	for _, arg := range args[1:] {
		if strings.Compare(arg, "-c") == 0 {
			desiredFlags = append(desiredFlags, arg)
		} else if strings.Compare(arg, "-l") == 0 { // Assuming file contains an extension
			desiredFlags = append(desiredFlags, arg)
		} else if strings.Compare(arg, "-w") == 0 {
			desiredFlags = append(desiredFlags, arg)
		} else if strings.Compare(arg, "-m") == 0 {
			desiredFlags = append(desiredFlags, arg)
		} else if strings.Contains(arg, ".") {
			fileToOpen = arg
		} else {
			return nil, "", fmt.Errorf("ccwc: illegal option -- %s\nusage: ./ccwc [-Lclmw] [file ...]", arg)
		}
	}

	return desiredFlags, fileToOpen, nil
}

// processFile opens the file and prints the desired field for each line.
func processFile(fileToOpen string, desiredFlags []string) error {
	

	if contains(desiredFlags, "-l") {
		lineCount, err := countLines(fileToOpen)
		if err != nil {
			return err
		}
		fmt.Printf("%d ", lineCount)
	}

	if contains(desiredFlags, "-w") {
		wordCount, err := countWords(fileToOpen)
		if err != nil {
			return err
		}
		fmt.Printf("%d ", wordCount)
	}

	if contains(desiredFlags, "-c") {
		charCount, err := countChars(fileToOpen)
		if err != nil {
			return err
		}
		fmt.Printf("%d ", charCount)
	}

	if !contains(desiredFlags, "-l") && !contains(desiredFlags, "-w") && !contains(desiredFlags, "-c") {
		lineCount, err := countLines(fileToOpen)
		if err != nil {
			return err
		}
		wordCount, err := countWords(fileToOpen)
		if err != nil {
			return err
		}
		charCount, err := countChars(fileToOpen)
		if err != nil {
			return err
		}
		fmt.Printf("%d %d %d ", lineCount, wordCount, charCount)
	}

	fmt.Printf("%s\n", fileToOpen)

	return nil
}
