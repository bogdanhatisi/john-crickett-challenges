package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a file name")
		return
	}
	fileName := args[2]
	if strings.Compare(args[1], "-c") == 0 {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	
		// Get the size in bytes
		size := fileInfo.Size()
	
		fmt.Printf("%d %s\n",size,fileName)
	}
}