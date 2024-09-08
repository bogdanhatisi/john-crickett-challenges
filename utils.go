package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertStringListToInt(stringList []string) []int {

	var intSlice []int
	for _, str := range stringList {

		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to int", err)
			continue
		}

		intSlice = append(intSlice, num)

	}

	return intSlice
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true // Target found
		}
	}
	return false // Target not found
}

func getMaxFieldNumber(desiredFields []int) int {
	max := 0
	for _, field := range desiredFields {
		if field > max {
			max = field
		}
	}
	return max
}

// parseArguments parses the command-line arguments and returns the desired field and the file to open.
func parseArguments(args []string) ([]int, string, string, error) {
	reNumber := regexp.MustCompile(`\d+`)
	var desiredFields []int
	var fileToOpen string = "-"
	var delimiter string = "\t"

	for _, arg := range args[1:] {
		if strings.Contains(arg, "-f") {
			fieldStr := reNumber.FindAllString(arg, -1)
			desiredFields = convertStringListToInt(fieldStr)
		} else if strings.Contains(arg, ".") { // Assuming file contains an extension
			fileToOpen = arg
		} else if strings.Contains(arg, "-d") {
			if len(arg) > 1 {
				delimiter = string(arg[len(arg)-1])
			} else {
				return nil, "", "", fmt.Errorf("delimiter flag -d provided without a delimiter character")
			}
		} else if strings.Compare(arg, "-") == 0 {
			fileToOpen = arg
		} else {
			return nil, "", "", fmt.Errorf("invalid argument: %s", arg)
		}
	}

	if len(desiredFields) == 0 || fileToOpen == "" {
		return nil, "", "", fmt.Errorf("please provide a valid column number and at least one argument")
	}

	return desiredFields, fileToOpen, delimiter, nil
}

// processFile opens the file and prints the desired field for each line.
func processFile(fileToOpen string, desiredFields []int, delimiter string) error {
	var maxFieldNumber int = getMaxFieldNumber(desiredFields)
	var scanner *bufio.Scanner
	if fileToOpen != "-" {

		file, err := os.Open(fileToOpen)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, delimiter)

		for index, field := range fields {

			if contains(desiredFields, index+1) {

				// Print the selected field
				fmt.Print(field)
				if index+1 != maxFieldNumber {
					fmt.Print(delimiter)
				}
			}

		}
		fmt.Println()
		// Ensure the desired field index is within range

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	return nil
}
