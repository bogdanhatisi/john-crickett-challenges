package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// parseArguments parses the command-line arguments and returns the desired field and the file to open.
func parseArguments(args []string) (int, string, error) {
	reNumber := regexp.MustCompile(`\d+`)
	var desiredField int
	var fileToOpen string

	for _, arg := range args[1:] {
		if strings.Contains(arg, "-f") {
			fieldStr := reNumber.FindString(arg)
			field, err := strconv.Atoi(fieldStr)
			if err != nil {
				return 0, "", fmt.Errorf("invalid field number: %v", err)
			}
			desiredField = field
		} else if strings.Contains(arg, ".") { // Assuming file contains an extension
			fileToOpen = arg
		} else {
			return 0, "", fmt.Errorf("invalid argument: %s", arg)
		}
	}

	if desiredField == 0 || fileToOpen == "" {
		return 0, "", fmt.Errorf("please provide a valid column number and at least one argument")
	}

	return desiredField, fileToOpen, nil
}

// processFile opens the file and prints the desired field for each line.
func processFile(fileToOpen string, desiredField int) error {
	file, err := os.Open(fileToOpen)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		// Ensure the desired field index is within range
		if desiredField > len(fields) {
			fmt.Printf("Field %d out of range for line: %s\n", desiredField, line)
			continue
		}

		// Print the selected field
		fmt.Printf("Field %d: %s\n", desiredField, fields[desiredField-1])
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	return nil
}