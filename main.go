package main

import (
	"os"
	"fmt"
	"ascii-art/ascii")


func main() {
	/*
	Reads the user input from 0s.Args and sends it to the render 
	function which handles the ASCII art generation
	*/
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		fmt.Println("Usage: go run . <input string> | cat e")
		return
	}

	input := os.Args[1]

	bannerLines, err := ascii.ReadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	asciiMap := ascii.BuildAsciiMap(bannerLines)

	result := ascii.PrintAscii(input, asciiMap)

	fmt.Print(result)
}