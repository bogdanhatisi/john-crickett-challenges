package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	desiredFlags, fileName, err := parseArguments(args)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = processFile(fileName, desiredFlags)

	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

}
