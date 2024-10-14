package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	// I am defining the flags here
	filename := flag.String("f", "", "filename")
	text := flag.String("d", "", "text to be saved in files")

	flag.Parse()

	// check if a filename given is empty string or empty text is provided 
	if *filename == "" || *text == "" {
		fmt.Printf("Please provide values for -f and -d flags")
		return
	}

	// writing text to the file
	err := os.WriteFile(*filename, []byte(*text), 0754)
	if err != nil {
		fmt.Printf("There was an error writing to this file: %v\n", err)
		return
	}

	fmt.Printf("Text written to file: %s\n", *filename)
}