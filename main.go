package main

import (
	"fmt"
	//"os"
	"ascii-art/ascii"
)

func main() {
	/*
	This function reads the user input from os.Args and sends it to the Render
	function which handles the ASCII art generation.
	*/
// 	if len(os.Args) != 2 {
// 		fmt.Println("Error: Invalid number of arguments")
// 		fmt.Println("Usage: go run . \"text\"")
// 		return
// 	}

// 	input := os.Args[1]

// 	ascii.ReadBanner(input, "standard.txt")
// }

// ----------------
	// I am using this section for testing the reading of the file to the CLI
	lines, err := ascii.ReadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
