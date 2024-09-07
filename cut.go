package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func main() {

	args := os.Args

	fmt.Println("All arguments:", args)

	var desiredFieldInteger int
	var fileToOpen string

	// Ensure that enough arguments are provided
	if len(args) < 3 {
		fmt.Println("Please provide a valid column number and at least one argument.")
		return
	}

	for argIndex, arg := range args {
		fmt.Println("Arg", argIndex, ":", arg)
		reNumber := regexp.MustCompile(`\d+`)
		if strings.Contains(arg, "-f") {
			desiredField := reNumber.FindString(arg)
			var err error
			desiredFieldInteger, err = strconv.Atoi(desiredField)
			if err != nil {
				fmt.Println("Invalid field number:", err)
				return
			}
		} else if strings.Contains(arg,".") {
			fileToOpen = arg
		} else{
			fmt.Println("Invalid argument:", arg)
			return
		}	
		
	}

	file, err := os.Open(fileToOpen)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read each line and split by tab
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		// Print the fields (or process them as needed)

		fmt.Printf("Field %d: %s\n", desiredFieldInteger, fields[desiredFieldInteger-1])

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
